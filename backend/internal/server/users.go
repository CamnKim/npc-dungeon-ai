package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"npc-dungeon-api/internal/database"
)

func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Guard clause to ensure only GET requests are allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	user, err := s.db.User().GetUserByID(id)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	jsonResp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Guard clause to ensure only POST requests are allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := &database.UserInsert{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	if user.Username == "" || user.Email == "" || user.Auth0Sub == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	insertedUser, err := s.db.User().CreateUser(user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(insertedUser)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) getUserByAuthHandler(w http.ResponseWriter, r *http.Request) {
	// Guard clause to ensure only GET requests are allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sub := r.PathValue("sub")
	user, err := s.db.User().GetUserByAuth0(sub)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found!", http.StatusNotFound)
		return
	}

	jsonResp, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
