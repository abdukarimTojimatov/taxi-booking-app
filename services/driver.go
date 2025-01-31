package services

import (
	"context"
	"errors"
	"taxi-booking-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverService interface {
	UpdateDriverLocation(driverID string, location models.GeoJSON) error
	GetAvailableDrivers() ([]models.Driver, error)
}

type driverService struct {
	collection *mongo.Collection
}

func NewDriverService(db *mongo.Database) DriverService {
	return &driverService{
		collection: db.Collection("drivers"),
	}
}

// UpdateDriverLocation - Haydovchining joylashuvini yangilash
func (ds *driverService) UpdateDriverLocation(driverID string, location models.GeoJSON) error {
	filter := bson.M{"user_id": driverID}
	update := bson.M{"$set": bson.M{"location": location}}

	_, err := ds.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("failed to update driver location")
	}

	return nil
}

// GetAvailableDrivers - Mavjud haydovchilarni olish
func (ds *driverService) GetAvailableDrivers() ([]models.Driver, error) {
	filter := bson.M{"is_available": true}

	cursor, err := ds.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.New("failed to fetch available drivers")
	}
	defer cursor.Close(context.TODO())

	var drivers []models.Driver
	if err := cursor.All(context.TODO(), &drivers); err != nil {
		return nil, errors.New("failed to decode drivers")
	}

	return drivers, nil
}