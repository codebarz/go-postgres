package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codebarz/go-postgres/internal/database"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	fmt.Println("Hello World!")

	godotenv.Load()

	port := os.Getenv("PORT")

	router := chi.NewRouter()

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("Postgres database URL not found in env")
	}

	conn, err := sql.Open("postgres", dbURL)

	apiHandlerConfig := apiConfig{
		DB: database.New(conn),
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	v1Routes := chi.NewRouter()

	v1Routes.Get("/health-check", handleReadiness)
	v1Routes.Get("/err", handleError)
	v1Routes.Post("/user", apiHandlerConfig.handleCreateUser)
	v1Routes.Get("/user", apiHandlerConfig.middlewareAuth(apiHandlerConfig.handleGetUserByApiKey))
	v1Routes.Post("/feed", apiHandlerConfig.middlewareAuth(apiHandlerConfig.handleCreateFeed))
	v1Routes.Get("/feed", apiHandlerConfig.handleGetFeeds)
	v1Routes.Post("/feed_follows", apiHandlerConfig.middlewareAuth(apiHandlerConfig.handleCreateFeedFollow))

	router.Mount("/v1", v1Routes)

	handler := c.Handler(router)

	server := &http.Server{
		Handler: handler,
		Addr:    ":" + port,
	}

	if port == "" {
		log.Fatal("PORT is not found as part of the environment variable")
	}

	log.Printf("Server staring on port %v", port)

	if err != nil {
		log.Printf("Error opening connection to DB %v:", err)
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started on port %v", port)
}
