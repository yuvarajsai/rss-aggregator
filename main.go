package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello world!")

	portString := os.Getenv("PORT") // gets the value of variable "PORT" from .env file
	if portString == "" {
		log.Fatal("PORT is not available in the environment.")
	}
	fmt.Println(portString)
}
