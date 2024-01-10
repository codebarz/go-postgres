package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	fmt.Println("Hello World!")

	godotenv.Load()

	port := os.Getenv("PORT")

	router := chi.NewRouter()

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

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started on port %v", port)
}
