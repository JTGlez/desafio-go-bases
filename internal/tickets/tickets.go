package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Ticket struct {
	Name        string
	Email       string
	Destination string
	FlightTime  time.Time
	Price       float64
}

var (
	ErrorMissingFile       = errors.New("missing tickets file")
	ErrorInvalidLineFormat = errors.New("invalid line format")
	ErrorInvalidID         = errors.New("invalid ID")
	ErrorInvalidFlightTime = errors.New("invalid flight time")
	ErrorInvalidPrice      = errors.New("invalid price")
	ErrorNoTicketsFound    = errors.New("no tickets found for the specified destination")
)

type TicketStore map[int]Ticket

func parseRecord(record []string) (Ticket, error) {
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

func LoadTicketData(filename string) (TicketStore, []error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, []error{ErrorMissingFile}
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, []error{err}
	}

	ts := make(TicketStore)
	var errors []error

	for i, record := range records {
		ticket, err := parseRecord(record)
		if err != nil {
			log.Printf("error parsing record %d: %v\n", i, err)
			errors = append(errors, err)
			continue
		}

		id, _ := strconv.Atoi(record[0])
		ts[id] = ticket
	}

	return ts, errors
}

func (ts TicketStore) Print() {

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

// ejemplo 2
func (ts TicketStore) CountByTimeOfDay() (madrugada, manana, tarde, noche int) {

	for _, ticket := range ts {
		hour := ticket.FlightTime.Hour()
		switch {
		case hour >= 0 && hour < 7:
			madrugada++
		case hour >= 7 && hour < 13:
			manana++
		case hour >= 13 && hour < 20:
			tarde++
		case hour >= 20 && hour <= 24:
			noche++
		}
	}
	return
}

// ejemplo 3
func AverageDestination(destination string, total int) (int error) {
	return nil
}
