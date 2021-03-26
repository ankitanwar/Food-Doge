package domain

import (
	"errors"
)

type Store struct {
	StoreID     string  `json:"storeID" bson:"_id,omitempty"`
	StoreName   string  `json:"storeName"`
	Street      string  `json:"street" bson:"street"`
	State       string  `json:"state" bson:"state"`
	Country     string  `json:"country" bson:"country"`
	PhoneNumber string  `json:"phoneNumber" bson:"phoneNumber"`
	Description string  `json:"description" bson:"description"`
	Pincode     int64   `json:"pincode" bson:"pincode"`
	StoreOwner  string  `json:"userID" bson:"userID"`
	Products    []Items `json:"products" bson:"products"`
}

func (details *Store) CheckForError() error {
	if details.StoreName == "" {
		return errors.New("Please Provide The valid store Name")
	} else if details.Country == "" {
		return errors.New("Please Enter The valid Country")
	} else if details.Street == "" {
		return errors.New("please Enter The valid street")
	} else if details.State == "" {
		return errors.New("Please Enter the valid state")
	} else if details.Pincode == 0 {
		return errors.New("please Enter the valid Pincode")
	} else if details.PhoneNumber == "" {
		return errors.New("Please Enter The Valid Phone Number")
	}
	return nil

}

type StoreLocation struct {
	Location                 string                   `json:"location" bson:"_id"`
	StoreLocationInformation StoreLocationInformation `json:"stores" bson:"stores"`
}
type StoreLocationInformation struct {
	StoreID     string `json:"storeID" bson:"storeID"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Street      string `json:"street" bson:"street"`
	State       string `json:"state" bson:"state"`
}

type StoreExplore struct {
	Location                 string                     `json:"location" bson:"_id"`
	StoreLocationInformation []StoreLocationInformation `json:"stores" bson:"stores"`
}
