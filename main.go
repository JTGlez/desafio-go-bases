package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	// LoadTicketData
	loader := tickets.CSVLoader{}
	store, err := loader.LoadTickets("tickets.csv")

	if len(err) != 0 {
		fmt.Println("Error loading ticket data:")
		for _, err := range err {
			fmt.Println(err)
		}
	}

	fmt.Println(store)

	// GetTotalTickets
	country := "Mexico"
	total, err := store.GetTotalTickets(country)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Total tickets to %s: %d\n", country, total)

	// CountByTimeOfDay
	earlyMorning, morning, afternoon, night := store.CountByTimeOfDay()
	fmt.Printf("Tickets by time of day:\n")
	fmt.Printf("Early Morning (0-6): %d\n", earlyMorning)
	fmt.Printf("Morning (7-12): %d\n", morning)
	fmt.Printf("Afternoon (13-19): %d\n", afternoon)
	fmt.Printf("Night (20-23): %d\n", night)

}
