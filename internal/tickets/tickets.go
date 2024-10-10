package tickets

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

type FlightTime struct {
	time.Time
}

func (ft *FlightTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	t, err := time.Parse("15:04", s)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrorInvalidFlightTime, err)
	}
	ft.Time = t
	return nil
}

type Ticket struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Destination string     `json:"destination"`
	FlightTime  FlightTime `json:"flight_time"`
	Price       float64    `json:"price"`
}

type TicketStore map[int]Ticket

// Slice de Personas que implementa sort.Interface
/* type PorEdad []Persona

func (p PorEdad) Len() int           { return len(p) }
func (p PorEdad) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PorEdad) Less(i, j int) bool { return p[i].Edad < p[j].Edad } */

type TicketLoader interface {
	LoadTickets(filename string) (TicketStore, []error)
}

var (
	ErrorMissingFile       = errors.New("missing tickets file")
	ErrorInvalidLineFormat = errors.New("invalid line format")
	ErrorInvalidID         = errors.New("invalid ID")
	ErrorInvalidFlightTime = errors.New("invalid flight time")
	ErrorInvalidPrice      = errors.New("invalid price")
	ErrorNoTicketsFound    = errors.New("no tickets found for the specified destination")
	ErrorStoreEmpty        = errors.New("ticket store is empty")
)

func (ts TicketStore) String() string {
	var sb strings.Builder
	keys := make([]int, 0, len(ts))
	for k := range ts {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		ticket := ts[k]
		sb.WriteString(fmt.Sprintf("ID: %d, Name: %s, Email: %s, Destination: %s, FlightTime: %s, Price: %.2f\n",
			ticket.ID, ticket.Name, ticket.Email, ticket.Destination, ticket.FlightTime, ticket.Price))
	}
	return sb.String()
}

func (ts TicketStore) GetTotalTickets(destination string) (total int, err error) {

	destination = strings.ToLower(destination)

	for _, ticket := range ts {
		if strings.ToLower(ticket.Destination) == destination {
			total++
		}
	}

	if total == 0 {
		err = ErrorNoTicketsFound
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
		case hour >= 20 && hour < 24:
			night++
		}
	}
	return
}

func (ts TicketStore) AverageDestination(destination string) (float64, error) {
	totalTickets := len(ts)
	if totalTickets == 0 {
		return 0, ErrorStoreEmpty
	}

	destinationTickets, err := ts.GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}

	percentage := (float64(destinationTickets) / float64(totalTickets)) * 100
	return percentage, nil
}
