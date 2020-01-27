package controller

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

// UserGrpcClient for connection auth grpc service
var UserGrpcClient *grpc.ClientConn

func init() {
	// get user service address
	USERADDRESS := "localhost:8080"
	conn, err := grpc.Dial(USERADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	UserGrpcClient = conn
}
