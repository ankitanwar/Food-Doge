package domain

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type ListALlProducts struct {
	AllProducts []Items `json:"products" bson:"products"`
}

type Items struct {
	ItemID      string `json:"itemID" bson:"itemID,omitempty"`
	ItemName    string `json:"itemName" bson:"itemName"`
	Cuisine     string `json:"cuisine" bson:"cuisine"`
	Description string `json:"description" bson:"description"`
	Vegetarain  bool   `json:"vegetarian" bson:"vegetarian"`
	Price       int64  `json:"price" bson:"price"`
}

func (details *Items) CheckItemDetails() error {
	if details.ItemName == "" {
		return errors.New("Item Name Cannot Be Empty")
	} else if details.Price == 0 {
		return errors.New("Please Enter A valid Price")
	}
	return nil
}

func (details *Items) GenerateUniqueID() (string, error) {
	id, := uuid.NewV4()
	stringID := id.String()
	return stringID, nil
}
