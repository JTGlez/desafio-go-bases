package tickets

import (
	"bufio"
	"encoding/json"
	"os"
)

type JSONLoader struct{}

func (j JSONLoader) LoadTickets(filename string) (TicketStore, []error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, []error{ErrorMissingFile}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var jsonData []byte

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		jsonData = append(jsonData, scanner.Bytes()...)
	}

	if err := scanner.Err(); err != nil {
		return nil, []error{err}
	}

	var tickets []Ticket
	err = json.Unmarshal(jsonData, &tickets)
	if err != nil {
		return nil, []error{err}
	}

	ts := make(TicketStore)

	for i, ticket := range tickets {
		ts[i+1] = ticket
	}

	return ts, nil
}

// ParseRecord Leftover: we can use unmarshal to directly to parse json data
func (j JSONLoader) ParseRecord(record []string) (Ticket, error) {
	return Ticket{}, nil
}
