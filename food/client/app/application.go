package application

import (
	"github.com/ankitanwar/Food-Doge/food/client/connect"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApllication() {
	UrlMapping()
	connect.ConnectServer()
	router.Run(":8070")

}
