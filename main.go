package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	// LoadTicketData from CSV
	csvLoader := tickets.CSVLoader{}
	csvStore, csvErr := csvLoader.LoadTickets("tickets.csv")

	if len(csvErr) != 0 {
		fmt.Println("Error loading ticket data from CSV:")
		for _, err := range csvErr {
			fmt.Println(err)
		}
	}

	fmt.Println("CSV Store:", csvStore)

	// LoadTicketData from JSON
	jsonLoader := tickets.JSONLoader{}
	jsonStore, jsonErr := jsonLoader.LoadTickets("tickets.json")

	if len(jsonErr) != 0 {
		fmt.Println("Error loading ticket data from JSON:")
		for _, err := range jsonErr {
			fmt.Println(err)
		}
	}
	fmt.Println("\n \n")
	fmt.Println("JSON Store:", jsonStore)

	// GetTotalTickets from JSON store
	country := "Finland"
	total, err := jsonStore.GetTotalTickets(country)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Total tickets to %s: %d\n", country, total)

	// CountByTimeOfDay from JSON store
	earlyMorning, morning, afternoon, night := jsonStore.CountByTimeOfDay()
	fmt.Printf("Tickets by time of day:\n")
	fmt.Printf("Early Morning (0-6): %d\n", earlyMorning)
	fmt.Printf("Morning (7-12): %d\n", morning)
	fmt.Printf("Afternoon (13-19): %d\n", afternoon)
	fmt.Printf("Night (20-23): %d\n", night)
}
