package application

import controllers "github.com/ankitanwar/Food-Doge/cart/controller"

func mapUrls() {
	router.POST("/user/cart/:itemID", controllers.AddToCart)
	router.DELETE("/user/cart/:itemID", controllers.RemoveFromCart)
	router.GET("/user/cart", controllers.ViewCart)
}
