package controllers

import (
	"context"
	"io"
	"log"
	"net/http"

	connect "github.com/ankitanwar/Food-Doge/stores/client/connect"
	storespb "github.com/ankitanwar/Food-Doge/stores/proto"
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
}

func (controller *foodControllerSturct) FilterFood(c *gin.Context) {
	storeID := c.Param("storeID")
	request := &storespb.FilterFoodRequest{}
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
	//need to implement authenctication
	userID := getUserID(c.Request)
	storeID := c.Param("storeID")
	itemDetails := &storespb.Food{}
	request := &storespb.AddItemsInStoreRequest{UserID: userID, StoreID: storeID}
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
	//need to implement authentication
	userID := getUserID(c.Request)
	storeID := c.Param("storeID")
	itemID := c.Param("itemID")
	request := &storespb.UpdateItemRequest{ItemID: itemID, StoreID: storeID, UserID: userID}
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
	userID := getUserID(c.Request)
	storeID := c.Param("storeID")
	itemID := c.Param("itemID")
	request := &storespb.DeleteItemRequest{ItemID: itemID, StoreID: storeID, UserID: userID}
	response, err := connect.Client.DeleteItem(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)

}

func (controller *foodControllerSturct) OrderFoodItem(c *gin.Context) {
	storeId := c.Param("storeID")
	itemID := c.Param("itemID")
	request := &storespb.OrderFoodRequest{
		StoreID: storeId,
		ItemID:  itemID,
	}
	response, err := connect.Client.OrderFood(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable To Order The food")
		return
	}
	c.JSON(http.StatusAccepted, response)
	return

}

func (controller *foodControllerSturct) GetItemDetails(c *gin.Context) {
	storeID := c.Param("storeID")
	itemID := c.Param("itemID")
	details := &storespb.GetItemDetailsRequest{
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
	req := &storespb.ViewParticularStoreRequest{
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
