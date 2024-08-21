package main

import (
	"flag"
	"fmt"
	"github/Doris-Mwito5/pcbook/pb"
	"github/Doris-Mwito5/pcbook/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	//get port from the cmd arg
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	// Create a new laptop server object
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	//register the laptop's server with grpc server
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	// Bind the server to the specified address and port
	//create an address string with the port
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	//listen for tcp connection on the server's address
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
		
	}
	//strat the server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}