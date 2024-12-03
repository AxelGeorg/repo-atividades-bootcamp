package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	pais := "Brazil"
	total := tickets.GetTotalTickets(pais)
	fmt.Println("\nQuantas pessoas viajam para o", pais, "é", total)

	periodo := "21:30"
	countHorario, err := tickets.GetCountByPeriod(periodo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Quantas pessoas viajam no Periodo das", periodo, "é", countHorario)
	}

	paisAvarage := "Brazil"
	porcentagem, errAvarage := tickets.AverageDestination(paisAvarage)
	if errAvarage != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Porcentagem de pessoas que viajam para o", paisAvarage, "é", porcentagem, "%")
	}
}
