package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roblkdeboer/url-shortener-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// var client *mongo.Client
var urlMappingsCollection *mongo.Collection

func ExpandURL(w http.ResponseWriter, r *http.Request) {
	// Extract the shortened URL parameter from the request
	vars := mux.Vars(r)
	shortenedURL := vars["shortened"]

	// Find the URLMapping document in the database
	var urlMapping models.URLMapping
	err := urlMappingsCollection.FindOne(context.Background(), bson.M{"_id": shortenedURL}).Decode(&urlMapping)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "URL not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error retrieving URL mapping", http.StatusInternalServerError)
		}
		return
	}

	redirectURL := "https://" + urlMapping.Original

	// Redirect to the original URL
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}