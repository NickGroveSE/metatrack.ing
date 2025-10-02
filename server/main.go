package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/NickGroveSE/metatrack.ing/models"
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
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data []models.OWHero
	data = append(data, models.OWHero{Name: "Ana", PickRate: 29.6, WinRate: 48.6})
	data = append(data, models.OWHero{Name: "Ashe", PickRate: 12, WinRate: 50.7})
	data = append(data, models.OWHero{Name: "Baptiste", PickRate: 9.6, WinRate: 46.2})

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
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/overwatch", owDataHandler)

	// Server configuration
	port := ":8080"
	log.Printf("Starting server on port %s", port)

	// Start server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
