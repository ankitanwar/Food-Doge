package controllers

import (
	"net/http"

	auth "github.com/ankitanwar/Food-Doge/middleware/auth"
	"github.com/ankitanwar/Food-Doge/stores/domain"
	"github.com/ankitanwar/Food-Doge/stores/services"
	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	order := &domain.PlaceOrder{}
	err := c.ShouldBindJSON(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error While Decoding The string")
	}
	storeID := c.Param("storeID")
	orderErr := services.PlaceOrder(storeID, order)
	if orderErr != nil {
		c.JSON(orderErr.Status, orderErr.Message)
		return
	}
	c.JSON(http.StatusAccepted, "Order Has Been Places Successfully")

}

func ViewOrder(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	storeID := c.Param("storeID")
	response, err := services.ViewOrders(storeID)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusAccepted, response)

}

func OrderCompleted(c *gin.Context) {
	if err := auth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	storeID := c.Param("storeID")
	orderID := c.Param("orderID")
	err := services.OrderCompleted(storeID, orderID)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusAccepted, "Order Has Been Completed Successfully")
}
