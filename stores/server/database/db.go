package db

import (
	"context"

	"github.com/ankitanwar/Food-Doge/stores/server/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StoreDB struct {
}

type ItemsDB struct {
}

func (storedb *StoreDB) SaveStore(details *domain.Store) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), details)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (StoreDB *StoreDB) DeleteStore(storeID primitive.ObjectID) error {
	filter := bson.M{"_id": storeID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (itemsDb *ItemsDB) DeleteItem(storeID primitive.ObjectID, userID, itemID string) error {
	filter := bson.M{"_id": storeID, "userID": userID}
	remove := bson.M{"$pull": bson.M{"products": bson.M{"itemID": itemID}}}
	_, err := collection.UpdateOne(context.Background(), filter, remove)
	if err != nil {
		return err
	}
	return nil
}

func (itemsDb *ItemsDB) AddItemToStore(storeID primitive.ObjectID, userID string, itemDetails *domain.Items) error {
	filter := bson.M{"_id": storeID, "userID": userID}
	add := bson.M{"$push": bson.M{"products": itemDetails}}
	_, err := collection.UpdateOne(context.Background(), filter, add)
	if err != nil {
		return err
	}
	return nil

}
