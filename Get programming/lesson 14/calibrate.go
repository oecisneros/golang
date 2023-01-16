package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

// go run calibrate.go
func main() {
	// https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
	rand.Seed(time.Now().UnixNano())

	var offset kelvin = 5
	sensor := calibrate(realSensor, offset)
	fmt.Println("Real sensor:", sensor())

	sensor = calibrate(fakeSensor, offset)
	for i := 0; i < 10; i++ {
		fmt.Println("Fake sensor:", sensor())
	}
}
