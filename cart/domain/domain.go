package domin

//Details : Slice for The Details Of The Items
type Details struct {
	UserID string `json:"userID" bson:"_id"`
	Detail []Item `json:"details" bson:"items"`
}

//Item : Item struct
type Item struct {
	ItemID      string `bson:"itemID" json:"itemID"`
	Description string `bson:"description" json:"description"`
	Price       int64  `bson:"price" json:"price"`
	Cuisine     string `bson:"cuisine" json:"cuisine"`
	StoreID     string `bson:"storeID" json:"storeID"`
}

type CheckoutCart struct {
	TotalPrice   int64  `json:"totalPrice"`
	DeliveryTime string `json:"deliveryTime"`
	Items        []Item `json:"items"`
	HouseNumber  string `json:"houseNo"`
	Street       string `json:"street"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
}
