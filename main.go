package main

import (
	"fmt"
	"log"
	"os"

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
}
