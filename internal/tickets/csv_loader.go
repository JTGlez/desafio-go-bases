package tickets

import (
	"encoding/csv"
	"fmt"
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

		}
	}(file)

	reader := csv.NewReader(file)
	ts := make(TicketStore)
	var errorChain []error

	for {
		record, err := reader.Read()

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			errorChain = append(errorChain, err)
			continue
		}

		if len(record) != 6 {
			errorChain = append(errorChain, fmt.Errorf("%w: record has %d fields, expected 6", ErrorInvalidLineFormat, len(record)))
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			errorChain = append(errorChain, fmt.Errorf("%w: %v", ErrorInvalidID, record[0]))
			continue
		}

		flightTime, err := time.Parse("15:04", record[4])
		if err != nil {
			errorChain = append(errorChain, fmt.Errorf("%w: %v", ErrorInvalidFlightTime, err))
			continue
		}

		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			errorChain = append(errorChain, fmt.Errorf("%w: %v", ErrorInvalidPrice, err))
			continue
		}

		ft := FlightTime{Time: flightTime}

		ticket := Ticket{
			ID:          id,
			Name:        record[1],
			Email:       record[2],
			Destination: record[3],
			FlightTime:  ft,
			Price:       price,
		}

		ts[id] = ticket
	}

	return ts, errorChain
}
