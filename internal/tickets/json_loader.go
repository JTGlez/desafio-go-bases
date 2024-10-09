package tickets

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type JSONLoader struct{}

func (j JSONLoader) LoadTickets(filename string) (TicketStore, []error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, []error{ErrorMissingFile}
	}

	var tickets []Ticket
	if err2 := json.Unmarshal(file, &tickets); err2 != nil {
		return nil, []error{err2}
	}

	ts := make(TicketStore)
	for _, ticket := range tickets {
		ts[ticket.ID] = ticket
	}

	return ts, nil
}

func (j JSONLoader) ParseRecord(record []string) (Ticket, error) {
	return Ticket{}, nil
}

func (t *Ticket) UnmarshalJSON(data []byte) error {
	type Alias Ticket
	aux := &struct {
		FlightTime string `json:"flight_time"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	flightTime, err := time.Parse("15:04", aux.FlightTime)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrorInvalidFlightTime, err)
	}
	t.FlightTime = flightTime

	return nil
}
