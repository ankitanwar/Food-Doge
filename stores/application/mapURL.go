package application

import "github.com/ankitanwar/Food-Doge/stores/controllers"

func mapUrls() {
	router.POST("/orders/:storeID", controllers.PlaceOrder)
	router.GET("/orders/store/:storeID", controllers.ViewOrder)
	router.DELETE("/orders/store/:storeID/:orderID", controllers.OrderCompleted)
}
