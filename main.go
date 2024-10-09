package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	// LoadTicketData
	store, err := tickets.LoadTicketData("tickets.csv")

	if len(err) != 0 {
		fmt.Println("Error loading ticket data:")
		for _, err := range err {
			fmt.Println(err)
		}
	}

	store.Print()

	// GetTotalTickets
	country := "Tibecuador"
	total, err := store.GetTotalTickets(country)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Total tickets to %s: %d\n", country, total)
	}

	// CountByTimeOfDay
	madrugada, manana, tarde, noche := store.CountByTimeOfDay()
	fmt.Printf("Tickets by time of day:\n")
	fmt.Printf("Madrugada (0-6): %d\n", madrugada)
	fmt.Printf("Mañana (7-12): %d\n", manana)
	fmt.Printf("Tarde (13-19): %d\n", tarde)
	fmt.Printf("Noche (20-23): %d\n", noche)
}
