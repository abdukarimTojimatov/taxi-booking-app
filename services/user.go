package services

import (
	"context"
	"errors"
	"taxi-booking-app/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(user *models.User) error
	LoginUser(email, password string) (string, error)
}

type userService struct {
	collection *mongo.Collection
}

func NewUserService(db *mongo.Database) UserService {
	return &userService{
		collection: db.Collection("users"),
	}
}

func (us *userService) CreateUser(user *models.User) error {
	_, err := us.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func (us *userService) LoginUser(email, password string) (string, error) {
	// Tizimga kirish logikasi
	return "jwt_token", nil
}