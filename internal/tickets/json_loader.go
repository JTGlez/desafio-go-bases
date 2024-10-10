package tickets

import (
	"encoding/json"
	"os"
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
