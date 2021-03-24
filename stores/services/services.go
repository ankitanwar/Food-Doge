package services

import (
	db "github.com/ankitanwar/Food-Doge/stores/database"
	"github.com/ankitanwar/Food-Doge/stores/domain"
	"github.com/ankitanwar/GoAPIUtils/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ViewOrders(storeID string) (*domain.ViewOrders, *errors.RestError) {
	if storeID == "" {
		return nil, errors.NewBadRequest("Invalid StoreID")
	}
	storeKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.NewInternalServerError("Unable To Decode The storeID")
	}
	orders := &domain.ViewOrders{}
	result := db.ViewOrders(storeKey)
	err = result.Decode(orders)
	if err != nil {
		return nil, errors.NewInternalServerError("Error While Decoding The Orders")
	}
	return orders, nil
}

func OrderCompleted(storeID, orderID string) *errors.RestError {
	storeKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return errors.NewInternalServerError("Unable To Decode The storeID")
	}
	err = db.DeleteOrder(storeKey, orderID)
	if err != nil {
		return errors.NewInternalServerError("Unable To Delete The Orde From The List")
	}
	return nil
}

func PlaceOrder(storeID string, order *domain.PlaceOrder) *errors.RestError {
	storeKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return errors.NewInternalServerError("Unable To Decode The storeID")
	}
	orderID, err := domain.GenerateOrderID()
	if err != nil {
		return errors.NewInternalServerError("Unable To Generate The order ID")
	}
	order.OrderID = orderID
	err = db.PlaceOrder(storeKey, order)
	if err != nil {
		return errors.NewInternalServerError("Unable To place The Order")
	}
	return nil
}
