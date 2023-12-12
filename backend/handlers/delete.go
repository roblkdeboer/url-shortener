package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteURL(w http.ResponseWriter, r *http.Request) {
	// Extract the shortened URL parameter from the request
	vars := mux.Vars(r)
	shortenedURL := vars["shortened"]

	// Delete the URLMapping document from the database
	_, err := urlMappingsCollection.DeleteOne(context.Background(), bson.M{"_id": shortenedURL})
	if err != nil {
		http.Error(w, "Error deleting URL mapping", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "URL deleted successfully"})
}