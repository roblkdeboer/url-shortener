package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/roblkdeboer/url-shortener-backend/models"
	"github.com/roblkdeboer/url-shortener-backend/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		URL string `json:"url"`
	}

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if the expected key ("url") is present
	if request.URL == "" {
		http.Error(w, "Missing field: url", http.StatusBadRequest)
		return
	}

	// Search the database for an existing URL mapping by its original URL
	var existingURLMapping models.URLMapping
	err := urlMappingsCollection.FindOne(context.Background(), bson.M{"original": request.URL}).Decode(&existingURLMapping)
	if err == nil {
		response := models.ResponseFormat{
			Key:      existingURLMapping.Shortened,
			LongURL:  existingURLMapping.Original,
			ShortURL: "http://localhost:8080/" + existingURLMapping.Shortened,
		}
		// If the URL already exists, respond with the existing mapping
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	} else if err != mongo.ErrNoDocuments {
		// Other error
		http.Error(w, "Error checking existing URL mapping", http.StatusInternalServerError)
		return
	}

	// Generate a unique short URL
	shortenedURL := utils.GenerateShortURL()

	// Create URLMapping document
	urlMapping := models.URLMapping{
		ID:        shortenedURL,
		Original:  request.URL,
		Shortened: shortenedURL,
		CreatedAt: time.Now(),
	}

	// Insert the document into the database
	_, err = urlMappingsCollection.InsertOne(context.Background(), urlMapping)
	if err != nil {
		http.Error(w, "Error storing URL mapping", http.StatusInternalServerError)
		return
	}

	// Return the shortened URL in the response
	response := models.ResponseFormat{
		Key:      shortenedURL,
		LongURL:  request.URL,
		ShortURL: "http://localhost:8080/" + shortenedURL,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}