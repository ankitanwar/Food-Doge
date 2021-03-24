package domain

import (
	"errors"

	uuid "github.com/uuid"
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
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	stringID := id.String()
	return stringID, nil
}
