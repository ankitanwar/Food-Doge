package controllers

import (
	"context"
	"io"
	"log"
	"net/http"

	connect "github.com/ankitanwar/Food-Doge/food/client/connect"
	foodpb "github.com/ankitanwar/Food-Doge/food/proto"
	auth "github.com/ankitanwar/Food-Doge/middleware/auth"
	orders "github.com/ankitanwar/Food-Doge/middleware/orders"
	"github.com/ankitanwar/Food-Doge/middleware/user"
	"github.com/gin-gonic/gin"
)

var FoodController foodControllerInterface = &foodControllerSturct{}

type foodControllerSturct struct {
}

type foodControllerInterface interface {
	AddNewItem(c *gin.Context)
	UpdateFoodDetails(c *gin.Context)
	DeleteFoodItem(c *gin.Context)
	FilterFood(c *gin.Context)
	OrderFoodItem(c *gin.Context)
	GetItemDetails(c *gin.Context)
	GetAllItems(c *gin.Context)
	CheckOut(c *gin.Context)
}

func (controller *foodControllerSturct) FilterFood(c *gin.Context) {
	storeID := c.Param("storeID")
	request := &foodpb.FilterFoodRequest{}
	err := c.ShouldBindJSON(request)
	request.StoreID = storeID
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error While Fetching The Filter Details")
		return
	}
	response, err := connect.Client.FilterDish(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error While Filtering The Stores")
		return
	}
	for {
		receivedDetails, err := response.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error while fetching the filtered details")
			return
		} else {
			c.JSON(http.StatusOK, receivedDetails)
		}
	}

}

func (controller *foodControllerSturct) AddNewItem(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	userID := getUserID(c.Request)
	storeID := c.Param("storeID")
	itemDetails := &foodpb.Food{}
	request := &foodpb.AddItemsInStoreRequest{UserID: userID, StoreID: storeID}
	err := c.ShouldBindJSON(itemDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	request.Details = itemDetails
	response, err := connect.Client.NewItem(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)

}

func (controller *foodControllerSturct) UpdateFoodDetails(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	userID := getUserID(c.Request)
	storeID := c.Param("storeID")
	itemID := c.Param("itemID")
	request := &foodpb.UpdateItemRequest{ItemID: itemID, StoreID: storeID, UserID: userID}
	err := c.ShouldBindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	response, err := connect.Client.UpdateItemDetail(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (controller *foodControllerSturct) DeleteFoodItem(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	userID := getUserID(c.Request)
	storeID := c.Param("storeID")
	itemID := c.Param("itemID")
	request := &foodpb.DeleteItemRequest{ItemID: itemID, StoreID: storeID, UserID: userID}
	response, err := connect.Client.DeleteItem(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)

}

func (controller *foodControllerSturct) OrderFoodItem(c *gin.Context) {
	addressID := c.Param("addressID")
	address, addressErr := user.GetUserAddress.GetAddress(c.Request, addressID)
	if addressErr != nil {
		c.JSON(addressErr.Status, addressErr.Message)
	}
	storeId := c.Param("storeID")
	itemID := c.Param("itemID")
	request := &foodpb.OrderFoodRequest{
		StoreID: storeId,
		ItemID:  itemID,
	}
	response, err := connect.Client.OrderFood(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable To Order The food")
		return
	}
	response.CustomerHouseNo = address.HouseNumber
	response.CustomerPhoneNumber = address.Phone
	response.CustomerStreet = address.Street
	response.CutomerState = address.State
	response.Pincode = address.Pincode
	informStoreErr := orders.PlaceOrders(c.Request, storeId, response)
	if informStoreErr != nil {
		c.JSON(informStoreErr.Status, informStoreErr.Message)
		return
	}
	c.JSON(http.StatusAccepted, response)
	return

}

func (controller *foodControllerSturct) GetItemDetails(c *gin.Context) {
	storeID := c.Param("storeID")
	itemID := c.Param("itemID")
	details := &foodpb.GetItemDetailsRequest{
		StoreID: storeID,
		ItemID:  itemID,
	}
	result, err := connect.Client.GetItemDetail(context.Background(), details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, result)

}

func (controller *foodControllerSturct) GetAllItems(c *gin.Context) {
	storeID := c.Param("storeID")
	if storeID == "" {
		c.JSON(http.StatusBadRequest, "Please Enter The Valid Store ID")
		return
	}
	req := &foodpb.ViewParticularStoreRequest{
		StoreID: storeID,
	}
	response, err := connect.Client.ViewItemOfStore(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable To Fetch The Details")
		return
	}
	for {
		details, err := response.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Unable To The Details Of The Item", err)
		}
		c.JSON(http.StatusAccepted, details)
	}

}
func (controller *foodControllerSturct) CheckOut(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	storeID := c.Param("storeID")
	itemID := c.Param("itemID")
	request := &foodpb.CheckOutRequest{ItemID: itemID, StoreID: storeID}
	response, err := connect.Client.Checkout(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable To Checkout The Item")
	}
	c.JSON(http.StatusAccepted, response)
	return

}
