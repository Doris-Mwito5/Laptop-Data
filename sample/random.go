package sample

import (
	"github/Doris-Mwito5/pcbook/pb"
	"math/rand"

	"github.com/google/uuid"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

//returns a random CPU brand by selecting one from a set of strings 
//randomCPUBrand func callas the randomstring fuction taht ahs two string args
func randomCPUBrand() string {
    return randomStringFromSet("Intel", "AMD")
}

//the function takes in variadic parameter that allows you to pass in a varaiable number of string arguments
func randomStringFromSet(a ...string) string {
	//calculates the no. of strings
    n := len(a)
	// no strings it returns an empty string
    if n == 0 {
        return ""
    }
	//otherwise it picks a random string from the set and returns  
    return a[rand.Intn(n)]
}

func randomCPUName(brand string) string{
	if brand == "Intel" {
		return randomStringFromSet(
            "Xeon E-2286M",
            "Core i9-9980HK",
            "Core i7-9750H",
            "Core i5-9400F",
            "Core i3-1005G1",
        )
    }

    return randomStringFromSet(
        "Ryzen 7 PRO 2700U",
        "Ryzen 5 PRO 3500U",
        "Ryzen 3 PRO 3200GE",
    )
	}

func randomint(min int, max int) int {
	return min + rand.Int()%(max-min+1)
}

func randomfloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomBool() bool { 
	return rand.Intn(2) == 1

}

func randomGPUBrand() string{
	return randomStringFromSet("Nvidia", "AMD")

}

func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
            "RTX 2060",
            "RTX 2070",
            "GTX 1660-Ti",
            "GTX 1070",
        )
    }

    return randomStringFromSet(
        "RX 590",
        "RX 580",
        "RX 5700-XT",
        "RX Vega-56",
    )
}

func randomfloat32(min, max float32) float32 {
    return min + rand.Float32()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
    if rand.Intn(2) == 1 {
        return pb.Screen_IPS
    }
    return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
    height := randomint(1080, 4320)
    width := height * 16 / 9

    resolution := &pb.Screen_Resolution{
        Width:  uint64(width),
        Height: uint64(height),
    }
    return resolution
}

func randomID() string{
	return uuid.New().String()
}

func randomLaptopBrand() string {
    return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
    switch brand {
    case "Apple":
        return randomStringFromSet("Macbook Air", "Macbook Pro")
    case "Dell":
        return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
    default:
        return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
    }
}
