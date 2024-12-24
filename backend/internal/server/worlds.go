package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"npc-dungeon-api/internal/database"
)

func (s *Server) getWorldByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// Get user from request header
	user := r.Header.Get("User")
	if user == "" {
		http.Error(w, fmt.Sprintf("Failed to get world with ID: %s", id), http.StatusUnauthorized)
		return
	}

	// Get world from db
	world, err := s.db.World().GetWorldById(id)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Failed to get world with ID: %s", id), http.StatusInternalServerError)
		return
	}
	if world == nil {
		http.Error(w, fmt.Sprintf("No world with ID: %s", id), http.StatusNotFound)
		return
	}

	// Check if user is owner of world
	if world.CreatedBy != user {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	jsonResp, err := json.Marshal(world)
	if err != nil {
		http.Error(w, "Failed to marshal json", http.StatusInternalServerError)
		return
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

	// decode and validate req body
	world := &database.WorldInsert{}
	if err := json.NewDecoder(r.Body).Decode(world); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// insert
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

func (s *Server) updateWorldHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// get user from request header
	user := r.Header.Get("User")
	if user == "" {
		http.Error(w, fmt.Sprintf("Failed to get world with ID: %s", id), http.StatusUnauthorized)
		return
	}

	// decode and validate request body
	body := &database.WorldUpdate{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// update database entry
	world, err := s.db.World().UpdateWorld(body, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, fmt.Sprintf("Failed to find world with ID: %s", id), http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to update world with ID: %s", id), http.StatusInternalServerError)
		return
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

func (s *Server) deleteWorldHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// get user from request header
	user := r.Header.Get("User")
	if user == "" {
		http.Error(w, fmt.Sprintf("Failed to get world with ID: %s", id), http.StatusUnauthorized)
		return
	}

	// Get world from db
	world, err := s.db.World().GetWorldById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get world with ID: %s", id), http.StatusInternalServerError)
		return
	}
	if world == nil {
		http.Error(w, fmt.Sprintf("No world with ID: %s", id), http.StatusNotFound)
		return
	}

	// Check if user is owner of world
	if world.CreatedBy != user {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	err = s.db.World().DeleteWorld(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete world with ID: %s", id), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) getUserWorldsHandler(w http.ResponseWriter, r *http.Request) {
	// Guard clause to ensure only GET requests are allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	uid := r.PathValue("uid")

	// get user from request header
	user := r.Header.Get("User")
	if user == "" {
		http.Error(w, fmt.Sprintf("Failed to get world with ID: %s", uid), http.StatusUnauthorized)
		return
	}

	// ensure ownership
	if uid != user {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// query
	worlds, err := s.db.World().GetAllWorldsByUserId(uid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get worlds owned by uid: %s", uid), http.StatusUnauthorized)
		return
	}

	if worlds == nil {
		http.Error(w, fmt.Sprintf("No worlds owned by uid: %s", uid), http.StatusNotFound)
		return
	}

	jsonResp, err := json.Marshal(worlds)
	if err != nil {
		http.Error(w, "Failed to marshal json", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
