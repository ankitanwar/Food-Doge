package db

import (
	"context"

	"github.com/ankitanwar/Food-Doge/food/server/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StoreDB struct {
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
