package tickets

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var testTickets TicketStore

func TestMain(m *testing.M) {
	loader := JSONLoader{}
	wd, err := os.Getwd()
	if err != nil {
		panic("failed to get working directory: " + err.Error())
	}
	jsonPath := filepath.Join(wd, "testDataset/tickets.json")

	var errs []error
	testTickets, errs = loader.LoadTickets(jsonPath)
	if len(errs) > 0 {
		panic("failed to load tickets: " + errs[0].Error())
	}

	m.Run()
}

func TestGetTotalTicketsSuite(t *testing.T) {
	suite := []struct {
		name          string
		destination   string
		expectedTotal int
		expectError   bool
	}{
		{
			name:          "Destino con tickets",
			destination:   "China",
			expectedTotal: 3,
			expectError:   false,
		},
		{
			name:          "Destino sin tickets",
			destination:   "Chile",
			expectedTotal: 0,
			expectError:   true,
		},
	}

	for _, tc := range suite {
		t.Run(tc.name, func(t *testing.T) {
			total, err := testTickets.GetTotalTickets(tc.destination)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedTotal, total)
			}
		})
	}
}

func TestAverageDestinationSuite(t *testing.T) {
	suite := []struct {
		name            string
		destination     string
		expectedPercent float64
		expectError     bool
	}{
		{
			name:            "Destino con tickets",
			destination:     "China",
			expectedPercent: 30.0,
			expectError:     false,
		},
		{
			name:            "TicketStore vacío",
			destination:     "Argentina",
			expectedPercent: 0.0,
			expectError:     true,
		},
	}

	for _, tc := range suite {
		t.Run(tc.name, func(t *testing.T) {
			percentage, err := testTickets.AverageDestination(tc.destination)
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedPercent, percentage)
			}
		})
	}
}

func TestCountByTimeOfDaySuite(t *testing.T) {
	earlyMorning, morning, afternoon, night := testTickets.CountByTimeOfDay()

	require.Equal(t, 4, earlyMorning, "Expected 3 tickets in early morning")
	require.Equal(t, 1, morning, "Expected 1 ticket in morning")
	require.Equal(t, 2, afternoon, "Expected 2 tickets in afternoon")
	require.Equal(t, 3, night, "Expected 4 tickets at night")
}
