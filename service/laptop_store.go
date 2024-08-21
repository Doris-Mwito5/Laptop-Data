package service

import (
	"errors"
	"fmt"
	"github/Doris-Mwito5/pcbook/pb"
	"sync"

	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("record already exist")

type LaptopStore interface {
	//has a save function that will save the laptop to the store
	Save(laptop *pb.Laptop) error
	//finds a laptop by id
	Find(id string) (*pb.Laptop, error)
}

//InMemory struct stores the laptop in store i.e implementingthe interface method
type InMemoryLaptopStore struct {
	//to handle the mutiple concurrent requests to save the laptop
	mutex sync.RWMutex
	data map[string]*pb.Laptop
}

//function to return a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		//initialize the data inside the function
		data: make(map[string]*pb.Laptop),
	}
}

//implement the Save function
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	
	//Check if the laptop id exist in the map or not
	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}
	//saving in the store using deep copy
	other := &pb.Laptop{}
	//cally copier to deepcopy the laptop object to other objects
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}
	store.data[other.Id] = other

	return nil
}

func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return other, nil
}