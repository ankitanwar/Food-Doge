package cotrollers

import (
	"net/http"

	"github.com/ankitanwar/Food-Doge/cart/services"
	auth "github.com/ankitanwar/Food-Doge/middleware/auth"
	user "github.com/ankitanwar/Food-Doge/middleware/user"
	"github.com/gin-gonic/gin"
)

func getCallerID(request *http.Request) string {
	userID := request.Header.Get("X-Caller-ID")
	return userID
}

//AddToCart : To add the given item into the cart
func AddToCart(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	userID := getCallerID(c.Request)
	itemID := c.Param("itemID")
	storeID := c.Param("storeID")
	err := services.AddToCart(userID, storeID, itemID)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusAccepted, "Item has been added to the cart successfully")
	return

}

//RemoveFromCart : To remove the given item from the cart
func RemoveFromCart(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	userID := getCallerID(c.Request)
	itemID := c.Param("itemID")
	err := services.RemoveFromCart(userID, itemID)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusAccepted, "Item has been removed successfully")
	return
}

//ViewCart : To view the cart of the particular user
func ViewCart(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	userID := getCallerID(c.Request)
	itemInCart, err := services.ViewCart(userID)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusAccepted, itemInCart)
	return
}

//Checkout : To checkout all the items from the cart
func Checkout(c *gin.Context) {
	//No need To verify The request because request will get verified in address service
	addressID := c.Param("addressID")
	address, err := user.GetUserAddress.GetAddress(c.Request, addressID)
	if err != nil {
		c.JSON(err.Status, err.Message)
	}
	userID := getCallerID(c.Request)
	response, err := services.Checkout(c.Request, userID)
	if err != nil {
		c.JSON(err.Status, err.Message)
	}
	response.Country = address.Country
	response.Street = address.Street
	response.State = address.State
	response.Phone = address.Phone
	response.HouseNumber = address.HouseNumber
	c.JSON(http.StatusAccepted, response)
}
