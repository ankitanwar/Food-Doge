package main

import (
	"os"
	"os/signal"

	application "github.com/ankitanwar/Food-Doge/stores/client/app"
	server "github.com/ankitanwar/Food-Doge/stores/server/startServer"
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
