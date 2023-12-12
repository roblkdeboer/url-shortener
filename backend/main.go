package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/roblkdeboer/url-shortener-backend/db"
	"github.com/roblkdeboer/url-shortener-backend/handlers"
)

func main() {
	// Connect to MongoDB
	client, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Set up the Gorilla Mux router
	router := mux.NewRouter()

	// Initialize handlers with the database connection
	handlers.InitHandlers(client)

	// Wrap the router with CORS middleware
	corsHandler := cors.Default().Handler(router)

	// Define API endpoints
	router.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	router.HandleFunc("/{shortened}", handlers.ExpandURL).Methods("GET")
	router.HandleFunc("/delete/{shortened}", handlers.DeleteURL).Methods("DELETE")

	// Start the server
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}





