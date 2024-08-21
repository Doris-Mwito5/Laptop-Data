package service

import (
	"context"
	"errors"
	"github/Doris-Mwito5/pcbook/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	Store LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	// Get the laptop object from the request
	laptop := req.GetLaptop()
	log.Printf("Received a CreateLaptop request with ID: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// Check if the ID is valid
		if _, err := uuid.Parse(laptop.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not valid: %v", err)
		}
	} else {
		// If the client hasn't sent an ID, generate one on the server
		id, err := uuid.NewRandom()
		if err != nil {
			// Return an internal server error if ID generation fails
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		// Set the laptop.ID to the generated ID
		laptop.Id = id.String()
	}

	// Save the laptop to the in-memory store
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists  // Corrected this line
		}
		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}

	log.Printf("Saved laptop with ID: %s", laptop.Id)

	// Create a new response with the laptop.ID and return it
	resp := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return resp, nil
}
