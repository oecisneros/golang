package main

import (
	"fmt"
	"math/rand"
)

type km int
type kms int

type ticket struct {
	spaceline string
	days      int
	triptype  string
	price     int
}

func rnd(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func getSpaceline(option int) string {
	switch option {
	case 1:
		return "Space Adventures"
	case 2:
		return "SpaceX"
	default:
		return "Virgin Galactic"
	}
}

func getTrip(option int) string {
	if option == 1 {
		return "One-way"
	}
	return "Round-trip"
}

func calculateDays(distance km, speed kms) int {
	seconds := int(distance) / int(speed)
	hours := seconds / 3600
	return hours / 24
}

func calculatePrice(speed kms, trip string) int {
	price := 20 + int(speed)
	if trip == "One-way" {
		return price
	}
	return price * 2
}

func newTicket() ticket {
	distance := km(62100000)
	speed := kms(rnd(16, 30))

	spaceline := getSpaceline(rnd(1, 3))
	days := calculateDays(distance, speed)
	trip := getTrip(rnd(1, 2))
	price := calculatePrice(speed, trip)

	return ticket{spaceline, days, trip, price}
}

func createTickets() []ticket {
	tickets := []ticket{}
	for i := 0; i < 20; i++ {
		tickets = append(tickets, newTicket())
	}
	return tickets
}

func printTickets(tickets []ticket) {
	fmt.Printf("%-18v %4v %-10v %5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("========================================")
	for _, item := range tickets {
		fmt.Printf("%-18v %4v %-10v $%4v\n", item.spaceline, item.days, item.triptype, item.price)
	}
}

// go run capstone.go
func main() {
	printTickets(createTickets())
}
