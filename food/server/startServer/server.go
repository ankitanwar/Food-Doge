package server

import (
	"log"
	"net"

	storespb "github.com/ankitanwar/Food-Doge/food/proto"
	"github.com/ankitanwar/Food-Doge/food/server/services"
	"google.golang.org/grpc"
)

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		log.Fatalln("Unable To start the server")
		return
	}
	opts := []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	storespb.RegisterStoresServiceServer(srv, &services.FoodService{})
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalln("Unable To Listen")
		return
	}

}
