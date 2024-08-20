package sample

import (
	"github/Doris-Mwito5/pcbook/pb"

	"github.com/golang/protobuf/ptypes"
)

//function that returns a pointer to tye object keyboard
func NewKeyboard() *pb.Keyboard {
	//creating keyboard object
	keyboard := &pb.Keyboard {
		Layout: randomKeyboardLayout(),
		Backlight: randomBool(),
	}
	return keyboard
}
//function to generate random cpu

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	
	numbercores := randomint(2, 8)
	numberthreads := randomint(numbercores, 12)

	minghz := randomfloat64(2.8, 3.2)
	maxghz := randomfloat64(minghz, 5.0)

	cpu := &pb.CPU {
		Brand: brand,
		Name: name,
		NumberCores: uint64(numbercores),
		NumberThreads: uint64(numberthreads),
		MinGhz: minghz,
		MaxGhz: maxghz,
	}
	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minghz := randomfloat64(1.0, 1.5)
	maxghz := randomfloat64(minghz, 2.0)

	memory := &pb.Memory {
		Value: uint64(randomint(2, 6)),
		Unit: pb.Memory_GIGABYTE,

	}

	gpu := &pb.GPU{
		Brand: brand,
		Name: name,
		MinGhz: minghz,
		MaxGhz: maxghz,
		Memory: memory,

	} 
	return gpu
}

func NewRAM() *pb.Memory {
    memGB := randomint(4, 64)

    ram := &pb.Memory{
        Value: uint64(memGB),
        Unit:  pb.Memory_GIGABYTE,
    }

    return ram
}

func NewSSD() *pb.Storage {
    
    ssd := &pb.Storage{
        Driver: pb.Storage_SSD,
        Memory: &pb.Memory{
            Value: uint64(randomint(128, 1024)),
            Unit:  pb.Memory_GIGABYTE,
        },
    }

    return ssd
}

// NewHDD returns a new sample HDD
func NewHDD() *pb.Storage {

    hdd := &pb.Storage{
        Driver: pb.Storage_HDD,
        Memory: &pb.Memory{
            Value: uint64(randomint(1, 6)),
            Unit:  pb.Memory_TERABYTE,
        },
    }

    return hdd
}

func NewScreen() *pb.Screen {
    screen := &pb.Screen{
        SizeInch:   randomfloat32(13, 17),
        Resolution: randomScreenResolution(),
        Panel:      randomScreenPanel(),
        Multitouch: randomBool(),
    }

    return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
    name := randomLaptopName(brand)

    laptop := &pb.Laptop{
        Id:       randomID(),
        Brand:    brand,
        Name:     name,
        Cpu:      NewCPU(),
        Ram:      NewRAM(),
        Gpus:     []*pb.GPU{NewGPU()},
        Storage: []*pb.Storage{NewSSD(), NewHDD()},
        Screen:   NewScreen(),
        Keyboard: NewKeyboard(),
        Weight:   &pb.Laptop_WeightKgs{
            WeightKgs: randomfloat64(1.0, 3.0),
        },
        PriceUsd:    randomfloat64(1500, 3500),
        ReleaseYear: int64(randomint(2015, 2019)),
        UpdatedAt:   ptypes.TimestampNow(),
    }

    return laptop
}