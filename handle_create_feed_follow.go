package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/codebarz/go-postgres/internal/database"
	"github.com/google/uuid"
)

func (dbConfig *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error decoding parameters %v", err))
		return
	}

	feedFollow, err := dbConfig.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating feed follow %v", err))
		return
	}

	respondWithJSON(w, 201, databseFeedFollowToFeedFollow(feedFollow))
}
