package app

import "github.com/ankitanwar/Food-Doge/user/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUser)
	router.PATCH("/users", controllers.UpdateUser)
	router.DELETE("/users", controllers.DeleteUser)
	router.GET("/user/address", controllers.GetAddress)
	router.POST("/user/address", controllers.AddAddress)
	router.POST("/user/verify", controllers.VerifyUser)
	router.GET("/user/address/:addressID")
}
