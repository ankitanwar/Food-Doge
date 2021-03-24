package main

import (
	"os"
	"os/signal"

	application "github.com/ankitanwar/Food-Doge/food/client/app"
	server "github.com/ankitanwar/Food-Doge/food/server/startServer"
)

func main() {
	go func() {
		application.StartApllication()
	}()
	go func() {
		server.StartServer()
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
