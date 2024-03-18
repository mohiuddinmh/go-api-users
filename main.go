package main

import (
	"fmt"
	"log"
	"net/http"

	"lets-go/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handlers.GetUsers)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
