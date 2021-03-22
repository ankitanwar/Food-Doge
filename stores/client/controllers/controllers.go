package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	connect "github.com/ankitanwar/Food-Doge/stores/client/connect"
	storespb "github.com/ankitanwar/Food-Doge/stores/proto"
	"github.com/gin-gonic/gin"
)

var (
	FoodController  foodControllerInterface  = &foodControllerSturct{}
	StoreController storeControllerInterface = &storeControllerStrcut{}
)

type foodControllerInterface interface {
	AddNewItem(c *gin.Context)
	OrderFood(c *gin.Context)
	UpdateFoodDetails(c *gin.Context)
	DeleteFoodItem(c *gin.Context)
}

type storeControllerInterface interface {
	CreateNewStore(c *gin.Context)
	ShowStores(c *gin.Context)
	UpdateStoreDetails(c *gin.Context)
	DeleteStore(c *gin.Context)
	FilterStores(c *gin.Context)
}

type foodControllerSturct struct {
}

type storeControllerStrcut struct {
}

func getUserID(req *http.Request) string {
	userID := req.Header.Get("userID")
	return userID
}

func (controller *storeControllerStrcut) CreateNewStore(c *gin.Context) {
	//need to implement authentication also
	userID := getUserID(c.Request)
	details := &storespb.CreateStoreRequest{}
	details.UserID = userID
	err := c.ShouldBindJSON(details)
	if err != nil {
		fmt.Println("the value of error si", err)
		c.JSON(http.StatusBadRequest, "Error While Fetching The Store Details")
	}
	fmt.Println("The value of userID is", details.UserID)
	response, err := connect.Client.CreateStore(context.Background(), details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, response)

}
func (controller *storeControllerStrcut) ShowStores(c *gin.Context) {
	location := c.Param("location")
	request := &storespb.ExploreOutletsRequest{
		Address: location,
	}
	respone, err := connect.Client.Explore(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	for {
		details, err := respone.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error While Fetching The store Details")
		} else {
			c.JSON(http.StatusOK, details)
		}
	}

}
func (controller *storeControllerStrcut) UpdateStoreDetails(c *gin.Context) {
	//need to implement authentication
	userID := getUserID(c.Request)
	storeID := c.Param("storeID")
	receivedDetails := &storespb.UpdateStoreRequest{}
	receivedDetails.StoreID = storeID
	receivedDetails.UpdatedDetails.UserID = userID
	err := c.ShouldBindJSON(receivedDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Unable To Fetch The Details")
		return
	}
	response, err := connect.Client.UpdateStoreDetails(context.Background(), receivedDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (controller *storeControllerStrcut) DeleteStore(c *gin.Context) {
	//need to implement authentication
	userID := getUserID(c.Request)
	storeID := c.Param("storeID")
	request := &storespb.DeleteStoreRequest{StoreID: storeID, UserID: userID}
	response, err := connect.Client.DeleteStore(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, response)

}
func (controller *storeControllerStrcut) FilterStores(c *gin.Context) {
	location := c.Param("location")
	request := &storespb.FilterOutletsRequest{}
	err := c.ShouldBindJSON(request)
	if request.Address == "" {
		request.Address = location
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error While Fetching The Filter Details")
		return
	}
	response, err := connect.Client.Filter(context.Background(), request)
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
	request := &storespb.AddItemsInStoreRequest{UserID: userID, StoreID: storeID}
	err := c.ShouldBindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	response, err := connect.Client.NewItem(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)

}

func (controller *foodControllerSturct) OrderFood(c *gin.Context) {

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
