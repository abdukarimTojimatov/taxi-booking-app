package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Driver struct {
	UserID      primitive.ObjectID `bson:"user_id"`
	Location    GeoJSON            `bson:"location"`
	IsAvailable bool               `bson:"is_available"`
}

type GeoJSON struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"` // [longitude, latitude]
}