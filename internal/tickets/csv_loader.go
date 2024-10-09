package tickets

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type CSVLoader struct{}

func (c CSVLoader) LoadTickets(filename string) (TicketStore, []error) {
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

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, []error{err}
	}

	ts := make(TicketStore)
	var errorChain []error

	for i, record := range records {
		ticket, err := c.ParseRecord(record)
		if err != nil {
			log.Printf("error parsing record %d: %v\n", i, err)
			errorChain = append(errorChain, err)
			continue
		}

		id, _ := strconv.Atoi(record[0])
		ts[id] = ticket
	}

	return ts, errorChain
}

func (c CSVLoader) ParseRecord(record []string) (Ticket, error) {
	if len(record) != 6 {
		return Ticket{}, fmt.Errorf("%w: record has %d fields, expected 6", ErrorInvalidLineFormat, len(record))
	}

	_, err := strconv.Atoi(record[0])
	if err != nil {
		return Ticket{}, fmt.Errorf("%w: %v", ErrorInvalidID, record[0])
	}

	flightTime, err := time.Parse("15:04", record[4])
	if err != nil {
		return Ticket{}, fmt.Errorf("%w: %v", ErrorInvalidFlightTime, err)
	}

	price, err := strconv.ParseFloat(record[5], 64)
	if err != nil {
		return Ticket{}, fmt.Errorf("%w: %v", ErrorInvalidPrice, err)
	}

	return Ticket{
		Name:        record[1],
		Email:       record[2],
		Destination: record[3],
		FlightTime:  flightTime,
		Price:       price,
	}, nil
}
