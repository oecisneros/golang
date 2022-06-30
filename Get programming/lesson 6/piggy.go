package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rnd(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func getCoin(option int) float64 {
	switch option {
	case 1:
		return 0.05
	case 2:
		return 0.10
	default:
		return 0.25
	}
}

// go run piggy.go
func main() {
	// https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
	rand.Seed(time.Now().UnixNano())

	balance := 0.0
	for balance < 20.0 {
		balance += getCoin(rnd(1, 3))
		fmt.Printf("%4.2f\n", balance)
	}
}
