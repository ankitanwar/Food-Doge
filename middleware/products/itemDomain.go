package product

//Details : Slice for The Details Of The Items
type Details struct {
	Detail Item `json:"details"`
}

//Item : Item struct
type Item struct {
	ItemID      string `bson:"itemID" json:"itemID"`
	Description string `bson:"description" json:"description"`
	Price       int64  `bson:"price" json:"price"`
	Cuisine     string `bson:"cuisine" json:"cuisine"`
	StoreID     string `bson:"storeID" json:"storeID"`
}

type Order struct {
	ItemName     string `json:"ItemName"`
	Price        int64  `json:"Price"`
	DeliveryTime string `json:"deliveryTime"`
}
