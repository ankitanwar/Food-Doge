package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Items struct {
	ItemID      uuid.UUID `json:"itemID" bson:"_id,omitempty"`
	ItemName    string    `json:"itemName" bson:"itemName"`
	Description string    `json:"description" bson:"description"`
	Vegetarain  bool      `json:"vegetarian" bson:"vegetarian"`
	Price       int64     `json:"price" bson:"price"`
}

func (details *Items) CheckItemDetails() error {
	if details.ItemName == "" {
		return errors.New("Item Name Cannot Be Empty")
	} else if details.Price == 0 {
		return errors.New("Please Enter A valid Price")
	}
	return nil
}

func (details *Items) GenerateUniqueID() (uuid.UUID, error) {
	id, err := uuid.New()
	if err != nil {
		return id, err
	}
	return id, nil
}
