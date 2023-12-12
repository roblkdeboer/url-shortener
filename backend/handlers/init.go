package handlers

import "go.mongodb.org/mongo-driver/mongo"

// InitHandlers initializes handlers with the MongoDB collection.
func InitHandlers(client *mongo.Client) {
   urlMappingsCollection = client.Database("urlshortener").Collection("urlmappings")
}