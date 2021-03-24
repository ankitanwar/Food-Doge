package connect

import (
	"fmt"

	storespb "github.com/ankitanwar/Food-Doge/food/proto"
	"google.golang.org/grpc"
)

var (
	Client storespb.StoresServiceClient
	CC     *grpc.ClientConn
)

//ConnectServer : To Connect To the gRPC server
func ConnectServer() {
	opts := grpc.WithInsecure()
	var err error
	CC, err = grpc.Dial("localhost:8082", opts)
	if err != nil {
		fmt.Println("Error while connection to the server", err.Error())
		panic(err)
	}
	Client = storespb.NewStoresServiceClient(CC)
	fmt.Println("Connection to Server is successfull")
}
