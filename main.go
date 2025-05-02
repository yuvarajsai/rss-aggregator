package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world!")

	godotenv.Load(".env") // imports

	portString := os.Getenv("PORT") // gets the value of variable "PORT" from .env file
	if portString == "" {
		log.Fatal("PORT is not available in the environment.")
	}
	fmt.Println("PORT:", portString)

	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server running on the port %v", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
