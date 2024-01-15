package main

import (
	"fmt"
	"net/http"

	"github.com/codebarz/go-postgres/internal/auth"
	"github.com/codebarz/go-postgres/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("auth error: %v", err))
			return
		}

		user, err := apiConfig.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("user not found with key: %v", apiKey))
			return
		}

		handler(w, r, user)
	}
}
