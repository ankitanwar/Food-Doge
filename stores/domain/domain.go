package domain

import "github.com/uuid"

type PlaceOrder struct {
	OrderID     string `json:"orderID" bson:"orderID"`
	HouseNumber string `json:"houseNo" json:"houseNumber"`
	Street      string `json:"street" json:"street"`
	State       string `json:"state" json:"state"`
	Phone       string `json:"phone" json:"phone"`
	Items       Order  `json:"order" bson:"order"`
}

type Order struct {
	ItemName string `json:"itemName" bson:"itemName"`
	Price    int64  `json:"price" bson:"price"`
}

type ViewOrders struct {
	StoreID string       `json:"storeID" bson:"_id"`
	Orders  []PlaceOrder `json:"orders" bson:"orders"`
}

func GenerateOrderID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	stringID := id.String()
	return stringID, nil
}
