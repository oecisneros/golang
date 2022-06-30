package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func rnd(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func getCoin(option int) int {
	switch option {
	case 1:
		return 5
	case 2:
		return 10
	default:
		return 25
	}
}

// go run piggy.go
func main() {
	// https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
	rand.Seed(time.Now().UnixNano())

	balance := 0
	for balance < 2000 {
		balance += getCoin(rnd(1, 3))
		fmt.Printf("dollars %3v, cents %3v\n", balance/100, balance%100)
	}

	//const distance = 24000000000000000000
	//fmt.Println("Distance 1", distance) //overflows

	distance2 := new(big.Int)
	distance2.SetString("24000000000000000000", 10)
	fmt.Println("Distance 2", distance2)
}
