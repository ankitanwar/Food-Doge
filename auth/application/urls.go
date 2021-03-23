package application

import controller "github.com/ankitanwar/Food-Doge/auth/controllers"

func mapURL() {
	router.POST("/login", controller.CreateAccessToken)
	router.GET("/validate", controller.ValidateAccessToken)
	router.DELETE("/logout", controller.RemoveAccessToken)
}
