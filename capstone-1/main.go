package main

import (
	"fmt"
	"math/rand"
)

const format = "%-20v %4v %-10v $ %4v\n"
const distanceToMars = 62100000
const minSpeed = 16
const secondsInDay = 3600 * 24

func main() {
	fmt.Printf("%-20v %4v %-10v %7v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Printf("============================================\n")
	for count := 10; count > 0; count--{
		spaceLine := getRandomSpaceLine()
		days, tripType, price := getRandomTrip()
		printLine(spaceLine, days, tripType, price)
	}
}

func getRandomTrip() (int, string, int) {
	randSpeed := rand.Intn(15)
	price := 36 + randSpeed
	multiplier := rand.Intn(2) + 1

	var tripType string
	switch multiplier {
	case 1:
		tripType = "One-way"
	case 2:
		tripType = "Round-trip"
	}

	speed := randSpeed + minSpeed

	days := distanceToMars / speed / secondsInDay
	return days, tripType, price * 2
}

func printLine(spaceLine string, days int, tripType string, price int) {
	fmt.Printf(format, spaceLine, days, tripType, price)
}

func getRandomSpaceLine() string {
	r := rand.Intn(3)
	var spaceLine string

	switch r {
	case 0:
		spaceLine = "Space Adventures"
	case 1:
		spaceLine =  "SpaceX"
	case 2:
		spaceLine = "Virgin Galactic"
	}
	return spaceLine
}
