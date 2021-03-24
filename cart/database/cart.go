package cartdatabase

import (
	"context"

	domain "github.com/ankitanwar/Food-Doge/cart/domain"
	product "github.com/ankitanwar/Food-Doge/middleware/Products"
	"github.com/ankitanwar/GoAPIUtils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//AddToCart : To add the item into the cart
func AddToCart(userID string, item product.Item) *errors.RestError {
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

//ViewCart : To view All the items in the cart
func ViewCart(userID string) (*domain.Details, error) {
	user := &domain.Details{}
	filter := bson.M{"_id": userID}
	err := collection.FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
