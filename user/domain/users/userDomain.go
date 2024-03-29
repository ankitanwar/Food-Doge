package users

import (
	"strings"

	"github.com/ankitanwar/GoAPIUtils/errors"
	uuid "github.com/satori/go.uuid"
)

//User : User and its values
type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
	PhoneNo     string `json:"phone"`
}

//Address : Address of the given user
type Address struct {
	UserID string        `bson:"_id" json:"userID"`
	List   []UserAddress `bson:"addresses" json:"addresses"`
}

//UserAddress : Address field for the user
type UserAddress struct {
	ID          string
	HouseNumber string `json:"houseNo"`
	Street      string `json:"street"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
	Pincode     int64  `json:"pincode"`
}

//Users : It will return the slices of users
type Users []User

//ValidateAddress : To validate the given aaddress
func (address *UserAddress) ValidateAddress() *errors.RestError {
	if len(address.HouseNumber) <= 0 {
		return errors.NewBadRequest("Enter the House Number")
	} else if len(address.Street) <= 0 {
		return errors.NewBadRequest("Please Enter The valid Street Number")
	} else if len(address.State) <= 0 {
		return errors.NewBadRequest("Enter the valid address")
	} else if len(address.Country) <= 0 {
		return errors.NewBadRequest("Enter the valid address")
	} else if len(address.Phone) > 10 || len(address.Phone) < 10 {
		return errors.NewBadRequest("Please Enter the valid phone number")
	} else if address.Pincode == 0 {
		return errors.NewBadRequest("Please Enter The Valid Pincode")
	}
	return nil
}

func (address *UserAddress) GenerateUniqueAddressID() (string, *errors.RestError) {
	id := uuid.NewV4()
	stringID := id.String()
	return stringID, nil
}

//Validate : To validate the users
func (user *User) ValidateDetails() *errors.RestError {
	if user.FirstName == "" {
		err := errors.NewBadRequest("Please Enter the First Name")
		return err
	}
	if user.LastName == "" {
		err := errors.NewBadRequest("Please Enter the Last Name")
		return err
	}
	if user.Email == "" {
		err := errors.NewBadRequest("Please enter the valid mail address")
		return err
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" || len(user.Password) < 5 {
		return errors.NewBadRequest("Please Enter the valid password")
	}
	if len(user.PhoneNo) != 0 {
		if len(user.PhoneNo) > 10 || len(user.PhoneNo) < 10 {
			return errors.NewBadRequest("Please Enter the valid phone Number")
		}
	}
	return nil
}
