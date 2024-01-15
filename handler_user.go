package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/codebarz/go-postgres/internal/database"
	"github.com/google/uuid"
)

func (dbConfig *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error decoding parameters %v", err))
		return
	}

	user, err := dbConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating %v", err))
		return
	}

	respondWithJSON(w, 201, databseUserToUser(user))
}

func (dbConfig *apiConfig) handleGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databseUserToUser(user))
}
