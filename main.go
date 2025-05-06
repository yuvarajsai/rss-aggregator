package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func fetchPort() string {
	godotenv.Load(".env") // imports

	portString := os.Getenv("PORT") // gets the value of variable "PORT" from .env file
	if portString == "" {
		log.Fatal("PORT is not available in the environment.")
	}
	fmt.Println("PORT:", portString)
	return portString
}

func main() {
	fmt.Println("Hello world!")
	portString := fetchPort()

	// Create main router
	router := chi.NewRouter()
	// Router CORS options
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	// GET {serverUrl}/v1/healthz
	v1Router.Get("/healthz", handlerReadiness)

	// GET {serverUrl}/v1/err
	v1Router.Get("/err", handlerError)

	// Mount v1 router on main router
	router.Mount("/v1", v1Router)
	// now GET http://localhost:8000/v1/healthz would return {} with 200 status code.
	// POST or any other method, would fail with 405 method not allowed error.

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server running on the port %v", portString)

	// Start the server and listen to requests
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
