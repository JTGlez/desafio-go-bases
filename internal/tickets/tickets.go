package tickets

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Ticket struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Destination string    `json:"destination"`
	FlightTime  time.Time `json:"flight_time"`
	Price       float64   `json:"price"`
}

type TicketStore map[int]Ticket

type TicketLoader interface {
	LoadTickets(filename string) (TicketStore, []error)
	ParseRecord(record []string) (Ticket, error)
}

var (
	ErrorMissingFile       = errors.New("missing tickets file")
	ErrorInvalidLineFormat = errors.New("invalid line format")
	ErrorInvalidID         = errors.New("invalid ID")
	ErrorInvalidFlightTime = errors.New("invalid flight time")
	ErrorInvalidPrice      = errors.New("invalid price")
	ErrorNoTicketsFound    = errors.New("no tickets found for the specified destination")
)

func (ts TicketStore) Stringer() {
	keys := make([]int, 0, len(ts))
	for k := range ts {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, ts[k])
	}
}

func (ts TicketStore) GetTotalTickets(destination string) (total int, err []error) {

	destination = strings.ToLower(destination)

	for _, ticket := range ts {
		if strings.ToLower(ticket.Destination) == destination {
			total++
		}
	}

	if total == 0 {
		err = []error{ErrorNoTicketsFound}
		return
	}

	return
}

func (ts TicketStore) CountByTimeOfDay() (earlyMorning, morning, afternoon, night int) {
	for _, ticket := range ts {
		hour := ticket.FlightTime.Hour()
		switch {
		case hour >= 0 && hour < 7:
			earlyMorning++
		case hour >= 7 && hour < 13:
			morning++
		case hour >= 13 && hour < 20:
			afternoon++
		case hour >= 20 && hour <= 24:
			night++
		}
	}
	return
}

// ejemplo 3
func AverageDestination(destination string, total int) (int error) {
	return nil
}
