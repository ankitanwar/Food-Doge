package cartdatabase

import (
	"context"
	"fmt"

	domain "github.com/ankitanwar/Food-Doge/cart/domain"
	"github.com/ankitanwar/GoAPIUtils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createNew(userID string, item domain.Item) error {
	t := &domain.User{}
	t.UserID = userID
	t.Items = append(t.Items, item)
	res, err := collection.InsertOne(context.Background(), t)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

//AddToCart : To add the item into the cart
func AddToCart(userID string, item domain.Item) *errors.RestError {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": userID}
	PushToCart := bson.M{"$push": bson.M{"items": item}}
	_, err := collection.UpdateOne(context.Background(), filter, PushToCart, opts)
	if err != nil {
		return errors.NewInternalServerError("Error While Adding Item into the cart")
	}
	return nil
}

//RemoveFromCart : To remove the item from the cart
func RemoveFromCart(userID, itemID string) error {
	filter := bson.M{"_id": userID}
	remove := bson.M{"$pull": bson.M{"items": bson.M{"itemID": itemID}}}
	_, err := collection.UpdateOne(context.Background(), filter, remove)
	if err != nil {
		return err
	}
	return nil
}

//Checkout : To checkout all the items in the cart
func Checkout(userID string) (*domain.User, error) {
	filter := bson.M{"_id": userID}
	user := &domain.User{}
	err := collection.FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

//ViewCart : To view All the items in the cart
func ViewCart(userID string) (*domain.User, error) {
	user := &domain.User{}
	filter := bson.M{"_id": userID}
	err := collection.FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}