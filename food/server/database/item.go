package db

import (
	"context"

	"github.com/ankitanwar/Food-Doge/food/server/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemsDB struct {
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

func (itemdb *ItemsDB) FilterItems(storeID primitive.ObjectID, price int64, cuisine, vegetarian, name string) (*mongo.Cursor, error) {
	filter := bson.D{{"_id", storeID}, {"products.itemName", name}, {"products.vegetarian", vegetarian}, {"products.cuisine", cuisine}}
	result, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (Itemsdb *ItemsDB) FetchAllItems(storeID primitive.ObjectID) *mongo.SingleResult {
	filter := bson.M{"_id": storeID}
	result := collection.FindOne(context.Background(), filter)
	return result
}

func (item *ItemsDB) GetItemDetail(storeID primitive.ObjectID, itemID string) *mongo.SingleResult {
	filter := bson.M{"_id": storeID, "products.itemID": itemID}
	result := collection.FindOne(context.Background(), filter)
	return result
}
