package db

import (
	"context"
	"fmt"

	"github.com/ankitanwar/Food-Doge/stores/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ViewOrders(storeID primitive.ObjectID) *mongo.SingleResult {
	filter := bson.M{"_id": storeID}
	result := collection.FindOne(context.Background(), filter)
	return result
}

func DeleteOrder(storeID primitive.ObjectID, orderID string) error {
	fmt.Println("The value of storeID and orderID is", storeID, orderID)
	filter := bson.M{"_id": storeID}
	remove := bson.M{"$pull": bson.M{"orders": bson.M{"orderID": orderID}}}
	_, err := collection.UpdateOne(context.Background(), filter, remove)
	if err != nil {
		return err
	}
	return nil
}

func PlaceOrder(storeID primitive.ObjectID, order *domain.PlaceOrder) error {
	filter := bson.M{"_id": storeID}
	opts := options.Update().SetUpsert(true)
	PushToCart := bson.M{"$push": bson.M{"orders": order}}
	_, err := collection.UpdateOne(context.Background(), filter, PushToCart, opts)
	if err != nil {
		return err
	}
	return nil
}
