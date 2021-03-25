package controllers

import (
	"context"
	"io"
	"log"
	"net/http"

	connect "github.com/ankitanwar/Food-Doge/food/client/connect"
	storespb "github.com/ankitanwar/Food-Doge/food/proto"
	auth "github.com/ankitanwar/Food-Doge/middleware/auth"
	"github.com/gin-gonic/gin"
)

var (
	StoreController storeControllerInterface = &storeControllerStrcut{}
)

type storeControllerInterface interface {
	CreateNewStore(c *gin.Context)
	ShowStores(c *gin.Context)
	UpdateStoreDetails(c *gin.Context)
	DeleteStore(c *gin.Context)
}

type storeControllerStrcut struct {
}

func getUserID(req *http.Request) string {
	userID := req.Header.Get("X-Caller-ID")
	return userID
}

func (controller *storeControllerStrcut) CreateNewStore(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	userID := getUserID(c.Request)
	details := &storespb.CreateStoreRequest{}
	details.UserID = userID
	err := c.ShouldBindJSON(details)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error While Fetching The Store Details")
	}
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
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
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
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
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
