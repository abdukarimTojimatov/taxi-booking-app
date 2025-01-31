package services

import (
	"context"
	"errors"
	"taxi-booking-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService interface {
	CreateOrder(order *models.Order) error
	GetOrderByID(orderID string) (*models.Order, error)
	UpdateOrderStatus(orderID string, status string) error
}

type orderService struct {
	collection *mongo.Collection
}

func NewOrderService(db *mongo.Database) OrderService {
	return &orderService{
		collection: db.Collection("orders"),
	}
}

// CreateOrder - Yangi buyurtma yaratish
func (os *orderService) CreateOrder(order *models.Order) error {
	order.ID = primitive.NewObjectID()
	_, err := os.collection.InsertOne(context.TODO(), order)
	if err != nil {
		return errors.New("failed to create order")
	}

	return nil
}

// GetOrderByID - Buyurtmani ID bo'yicha olish
func (os *orderService) GetOrderByID(orderID string) (*models.Order, error) {
	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}

	filter := bson.M{"_id": objID}
	var order models.Order
	err = os.collection.FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		return nil, errors.New("order not found")
	}

	return &order, nil
}
// UpdateOrderStatus - Buyurtma holatini yangilash
func (os *orderService) UpdateOrderStatus(orderID string, status string) error {
	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return errors.New("invalid order ID")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"status": status}}

	_, err = os.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("failed to update order status")
	}

	return nil
}
// UpdateOrderStatus - Buyurtma