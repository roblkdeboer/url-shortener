package models

import "time"

type URLMapping struct {
	ID        string    `json:"id" bson:"_id"`
	Original  string    `json:"original" bson:"original"`
	Shortened string    `json:"shortened" bson:"shortened"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}