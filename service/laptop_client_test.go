package service_test

import (
	"context"
	"github/Doris-Mwito5/pcbook/pb"
	"github/Doris-Mwito5/pcbook/sample"
	"github/Doris-Mwito5/pcbook/serializer"
	"github/Doris-Mwito5/pcbook/service"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TestClientCreateLaptop tests the CreateLaptop gRPC service method
func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopServer, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddress)

	// Create a new sample laptop
	laptop := sample.NewLaptop()

	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// Call the CreateLaptop method on the client
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	// Check if the laptop is stored on the server
	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// Check if the saved laptop is the same as the one sent
	requireSameLaptop(t, laptop, other)
}

// startTestLaptopServer starts a gRPC server for testing
func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	// Create a new laptop server
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register the laptop service server on the gRPC server
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	// Create a listener that will listen on a random available port
	listener, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)

	// Run the gRPC server in a separate goroutine
	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

// newTestLaptopClient creates a new gRPC client for testing
func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

// requireSameLaptop checks if two laptops are the same by comparing their JSON representations
func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
