package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"npc-dungeon-api/internal/database"
)

func (s *Server) getWorldByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Guard clause to ensure only GET requests are allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.PathValue("id")

	world, err := s.db.World().GetWorldById(id)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get world with ID: %s", id), http.StatusInternalServerError)
	}
	if world == nil {
		http.Error(w, fmt.Sprintf("No world with ID: %s", id), http.StatusNotFound)
	}

	jsonResp, err := json.Marshal(world)
	if err != nil {
		http.Error(w, "Failed to marshal json", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) createWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Guard clause to ensure only GET requests are allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	world := &database.WorldInsert{}
	if err := json.NewDecoder(r.Body).Decode(world); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	inserted, err := s.db.World().CreateWorld(world)
	if err != nil {
		http.Error(w, "Failed to insert", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(inserted)
	if err != nil {
		http.Error(w, "Failed to marshal json", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
