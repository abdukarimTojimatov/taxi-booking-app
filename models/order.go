package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	DriverID  primitive.ObjectID `bson:"driver_id"`
	From      GeoJSON             `bson:"location"`
	To        GeoJSON             `bson:"location"`
	Status    string             `bson:"status"` // "pending", "accepted", "completed"
	Price     float64            `bson:"price"`
	CreatedAt time.Time          `bson:"created_at"`
}