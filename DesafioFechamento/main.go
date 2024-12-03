package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	total := tickets.GetTotalTickets("Brazil")

	fmt.Println(total)
}
