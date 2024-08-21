package main

import (
	"context"
	"flag"
	"github/Doris-Mwito5/pcbook/pb"
	"github/Doris-Mwito5/pcbook/sample"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	//get the server address from the cmd
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	//log for dialling the serverAddress
	log.Printf("dial server %s", *serverAddress)

	//call grpc
	conn, err := grpc.NewClient(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
		
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("laptop already exists")
		} else {
			log.Fatal("cannot create laptop")
		}

		return
	}
	log.Printf("craeted laptop with id: %s", res.Id)

}