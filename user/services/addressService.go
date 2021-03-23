package services

import (
	addressDB "github.com/ankitanwar/Food-Doge/user/databasource/mongoDB"
	"github.com/ankitanwar/Food-Doge/user/domain/users"
	"github.com/ankitanwar/GoAPIUtils/errors"
)

var (
	AddresService addressServiceInterface = &addressServiceStruct{}
)

type addressServiceInterface interface {
	GetAddress(string) (*users.Address, *errors.RestError)
	AddAddress(string, *users.UserAddress) *errors.RestError
	RemoveAddress(string, string) *errors.RestError
	GetAddressWithID(string, string) (*users.UserAddress, *errors.RestError)
}

type addressServiceStruct struct{}

func (userAddress *addressServiceStruct) GetAddress(userID string) (*users.Address, *errors.RestError) {
	details, err := addressDB.GetUserAddress(userID)
	if err != nil {
		return nil, errors.NewInternalServerError("Error While fetching the address")
	}
	return details, nil
}

func (userAddress *addressServiceStruct) AddAddress(userID string, address *users.UserAddress) *errors.RestError {
	err := address.ValidateAddress()
	if err != nil {
		return err
	}
	id, err := address.GenerateUniqueAddressID()
	if err != nil {
		return err
	}
	address.ID = id
	saveErr := addressDB.AddAddress(userID, address)
	if saveErr != nil {
		return errors.NewBadRequest("Error While Saving The Address Into The Databse")
	}
	return nil

}

func (userAddress *addressServiceStruct) RemoveAddress(userID string, addressID string) *errors.RestError {
	err := addressDB.RemoveAddress(userID, addressID)
	if err != nil {
		return errors.NewInternalServerError("Unable To Delete The Address")
	}
	return nil
}

func (userAddress *addressServiceStruct) GetAddressWithID(userID, addressID string) (*users.UserAddress, *errors.RestError) {
	if addressID == "" {
		return nil, errors.NewBadRequest("Please Enter The valid Address ID")
	}
	details := addressDB.UserSpecificAddress(userID, addressID)
	if details == nil {
		return nil, errors.NewBadRequest("Given Address Not Found In The Database")
	}
	address := &users.UserAddress{}
	err := details.Decode(address)
	if err != nil {
		return nil, errors.NewInternalServerError("Error While Decoding The Address")
	}
	return address, nil
}
