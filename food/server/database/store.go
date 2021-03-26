package db

import (
	"context"

	"github.com/ankitanwar/Food-Doge/food/server/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (StoreDB *StoreDB) AddStoreInLocationSearch(details *domain.StoreLocation) error {
	filter := bson.M{"_id": details.Location}
	opts := options.Update().SetUpsert(true)
	add := bson.M{"$push": bson.M{"stores": details.StoreLocationInformation}}
	_, err := storeCollection.UpdateOne(context.Background(), filter, add, opts)
	if err != nil {
		return err
	}
	return nil

}

func (StoreDB *StoreDB) ExploreStore(state string) *mongo.SingleResult {
	filter := bson.M{"_id": state}
	result := storeCollection.FindOne(context.Background(), filter)
	return result
}
