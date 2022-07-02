package main

import (
	"fmt"
)

func main() {
	const distanceToCanis = 236000000000000000 // km
	const lightSpeed = 300000                  // km/s
	const secondsPerDay = 86400
	const lightYear = lightSpeed * secondsPerDay * 365 // km/s

	const totalYears = distanceToCanis

	fmt.Println("Canis Major Dwarf Galaxy is", totalYears, "light years away.")
}
