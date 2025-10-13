package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/NickGroveSE/metatrack.ing/backend/models"
	"github.com/NickGroveSE/metatrack.ing/backend/scrape"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
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

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4321")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()

	input := queryParams.Get("input")
	log.Println(input)
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
	// Register handler
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/overwatch", owDataHandler)

	// Server configuration
	port := ":8080"
	log.Printf("Starting server on port %s", port)

	// Start server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
