package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/NickGroveSE/metatrack.ing/backend/models"
	"github.com/NickGroveSE/metatrack.ing/backend/scrape"
	"golang.org/x/time/rate"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

// IPRateLimiter manages rate limiters for each IP
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	limiter := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
	// Start cleanup routine
	limiter.CleanupStale(5 * time.Minute)
	return limiter
}

func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()
	return limiter
}

// Cleanup old limiters periodically to prevent memory leaks
func (i *IPRateLimiter) CleanupStale(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			i.mu.Lock()
			i.ips = make(map[string]*rate.Limiter)
			i.mu.Unlock()
		}
	}()
}

// rateLimitMiddleware wraps handlers with rate limiting
func rateLimitMiddleware(limiter *IPRateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract IP address
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				ip = r.RemoteAddr // Fallback to full address
			}

			// Check rate limit
			l := limiter.GetLimiter(ip)
			if !l.Allow() {
				http.Error(w, "Too Many Requests - Please slow down", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Message:   "Server is running",
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func owDataHandler(w http.ResponseWriter, r *http.Request) {

	if origin := r.Header.Get("Origin"); origin == "http://localhost:4321" || origin == "https://metatrack.ing" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()

	input := queryParams.Get("input")
	owmap := queryParams.Get("map")
	region := queryParams.Get("region")
	role := queryParams.Get("role")
	queue := queryParams.Get("queue")
	rank := queryParams.Get("rank")

	var data []models.OWHero = scrape.Scrape(input, owmap, region, role, queue, rank)

	response := models.OWDataResponse{
		Status:    "success",
		Timestamp: time.Now(),
		Data:      data,
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func main() {
	// Create rate limiter: 10 requests per second per IP, burst of 20
	limiter := NewIPRateLimiter(10, 20)

	// Create a new mux
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/overwatch", owDataHandler)

	// Wrap mux with rate limiting middleware
	handler := rateLimitMiddleware(limiter)(mux)

	// Server configuration
	port := ":8080"
	log.Printf("Starting server on port %s with rate limiting enabled", port)

	// Start server
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
