package main

import (
	"fmt"
	"log"
	"net"
	"os"

	grpcService "github.com/Penetration-Platform-Go/Auth-Service/grpc"
	"github.com/Penetration-Platform-Go/Auth-Service/route"
	auth "github.com/Penetration-Platform-Go/gRPC-Files/Auth-Service"
	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"
)

func main() {
	var PORT = os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8001"
	}
	var port = flag.StringP("port", "p", PORT, "Define the port where service runs")
	var GRPCPORT = os.Getenv("GRPC_PORT")
	if len(GRPCPORT) == 0 {
		GRPCPORT = "8081"
	}
	var grpcPort = flag.StringP("grpc_port", "g", GRPCPORT, "Define the port where grpc service runs")
	flag.Parse()
	// start grpc server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *grpcPort))
		if err != nil {
			log.Fatalf("failed to listen grpc: %v", err)
		}
		log.Printf("Listening on: %s", *grpcPort)
		gs := grpc.NewServer()
		auth.RegisterAuthServer(gs, &grpcService.AuthService{})
		gs.Serve(lis)
	}()

	s := route.GetServer()
	s.Run(":" + *port)
}
