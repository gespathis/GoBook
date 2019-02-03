package main

import (
	"GoBook/RandTicketsToMars/providers"
	"fmt"
)

func main() {
	const earthToMars = 62100000 //km
	const secondsInDay = 86400
	days := 0
	ticketsToFind := 10
	var totalPrice float32

	fmt.Printf("\n")
	fmt.Printf("%-20v%-5v%-12v%v\n", "SpaceLine", "Days", "Trip type", "Price")
	fmt.Printf("========================================================================\n")

	for index := 0; index < ticketsToFind; index++ {

		tripType := "One-way"

		company, multiplier, speed, priceTrip, oneWay := providers.GetCompanyToGoMars()
		if speed > 0 {
			days = earthToMars / speed / secondsInDay
			totalPrice = float32(priceTrip) * multiplier

			if oneWay == 0 {
				totalPrice = totalPrice * 2
				tripType = "Round-trip"
			}
		}

		fmt.Printf("%-20v%-5v%-11v %4.2f\n", company, days, tripType, totalPrice)
	}
}
