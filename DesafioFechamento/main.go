package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	tickets.FillTicketList()

	country := "Brazil"
	total := tickets.GetTotalTickets(country)
	fmt.Println("\nHow many people are traveling to", country, "is", total)

	period := "21:30"
	countPeriod, err := tickets.GetCountByPeriod(period)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("How many people are traveling in the period of", period, "is", countPeriod)
	}

	countryAverage := "Brazil"
	percentage, errAverage := tickets.AverageDestination(countryAverage)
	if errAverage != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Percentage of people traveling to", countryAverage, "is", percentage, "%")
	}
}
