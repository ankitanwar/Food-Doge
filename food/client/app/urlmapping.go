package application

import "github.com/ankitanwar/Food-Doge/food/client/controllers"

func UrlMapping() {
	router.GET("/stores/:location", controllers.StoreController.ShowStores)
	router.GET("/store/:storeID/filter", controllers.FoodController.FilterFood)
	router.POST("/stores/newstore", controllers.StoreController.CreateNewStore)
	router.PATCH("/store/:storeID/update", controllers.StoreController.UpdateStoreDetails)
	router.DELETE("/store/:storeID/delete", controllers.StoreController.DeleteStore)
	router.POST("/food/:storeID", controllers.FoodController.AddNewItem)
	router.PATCH("/food/:storeID/:itemID/updateItem", controllers.FoodController.UpdateFoodDetails)
	router.DELETE("/food/:storeID/:itemID", controllers.FoodController.DeleteFoodItem)
	router.GET("/food/all/:storeID", controllers.FoodController.GetAllItems)
	router.GET("/itemDetail/:storeID/:itemID", controllers.FoodController.GetItemDetails)
	router.POST("/buy/food/:storeID/:itemID", controllers.FoodController.OrderFoodItem)
	router.POST("/checkout/:storeID/:itemID", controllers.FoodController.CheckOut)
}
