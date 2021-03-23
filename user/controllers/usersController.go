package controllers

import (
	"fmt"
	"net/http"

	auth "github.com/ankitanwar/Food-Doge/middleware/auth"
	"github.com/ankitanwar/Food-Doge/user/domain/users"
	"github.com/ankitanwar/Food-Doge/user/services"
	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/gin-gonic/gin"
)

func getUserid(request *http.Request) (string, *errors.RestError) {
	userID := request.Header.Get("X-Caller-ID")
	if userID == "" {
		return "", errors.NewBadRequest("Invalid User ID")
	}
	return userID, nil
}

//CreateUser : To create the user
func CreateUser(c *gin.Context) {
	var newUser users.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		err := errors.NewBadRequest("Invalid Request")
		c.JSON(err.Status, err)
		return
	}

	_, saverr := services.UserServices.CreateUser(newUser)
	if saverr != nil {
		c.JSON(saverr.Status, saverr)
		return
	}
	c.JSON(http.StatusAccepted, "User Has Been Created Successfully")
}

//GetUser : To get the user from the database
func GetUser(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	userid, userErr := getUserid(c.Request)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user, err := services.UserServices.GetUser(userid)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	response := &users.ReturnUserDetails{}
	response.ShowDetails(user)
	c.JSON(http.StatusOK, response)

}

//UpdateUser :To Update the value of particaular user
func UpdateUser(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	var user = users.User{}
	userid, userErr := getUserid(c.Request)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user.ID = userid
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	updatedUser, err := services.UserServices.UpdateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	response := &users.ReturnUserDetails{}
	response.ShowDetails(updatedUser)
	c.JSON(http.StatusOK, response)
}

//DeleteUser :To Delete the user with given id
func DeleteUser(c *gin.Context) {
	userid, userErr := getUserid(c.Request)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	if err := services.UserServices.DeleteUser(userid); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"Status": "User Deleted"})
}

//Login : to verify user email and password
func VerifyUser(c *gin.Context) {
	verifyUser := users.LoginRequest{}
	if err := c.ShouldBindJSON(&verifyUser); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user, err := services.UserServices.VerifyUser(verifyUser)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	response := &users.ReturnUserDetails{}
	response.ShowDetails(user)
	c.JSON(http.StatusOK, response)

}

//GetAddress : To Get the address of the given user
func GetAddress(c *gin.Context) {
	err := auth.AuthenticateRequest(c.Request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	userID, err := getUserid(c.Request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	address, err := services.AddresService.GetAddress(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusAccepted, address)
}

//AddAddress : To Save the address of the given user
func AddAddress(c *gin.Context) {
	err := auth.AuthenticateRequest(c.Request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	address := &users.UserAddress{}
	bindErr := c.ShouldBindJSON(address)
	fmt.Println("The value of bindErr is", bindErr)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, "Error while binding to the json")
		return
	}
	userID, err := getUserid(c.Request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	err = services.AddresService.AddAddress(userID, address)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusAccepted, "Address has been added successfully")
}

func GetAddressWithID(c *gin.Context) {
	err := auth.AuthenticateRequest(c.Request)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	addressID := c.Param("addressID")
	userID, err := getUserid(c.Request)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	response, err := services.AddresService.GetAddressWithID(userID, addressID)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusAccepted, response)
	return
}
