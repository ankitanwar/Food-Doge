package application

import "github.com/ankitanwar/Food-Doge/stores/client/controllers"

func UrlMapping() {
	router.GET("/stores/:location", controllers.StoreController.ShowStores)
	router.GET("/stores/:location/filter", controllers.StoreController.FilterStores)
	router.POST("/stores", controllers.StoreController.CreateNewStore)
	router.PATCH("/store/:storeID", controllers.StoreController.UpdateStoreDetails)
	router.DELETE("/store/:storeID", controllers.StoreController.DeleteStore)
	router.POST("/food/:storeID", controllers.FoodController.AddNewItem)
	router.PATCH("/food/:storeID/:itemID", controllers.FoodController.UpdateFoodDetails)
	router.DELETE("/food/:storeID/:itemID", controllers.FoodController.DeleteFoodItem)
}
