package tickets_test

import (
	"testing"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
)

func FillTicketList() {
	if tickets.GetCountTicketList() == 0 {
		tickets.FillTicketList()
	}
}

func TestGetTotalTickets(t *testing.T) {

	FillTicketList()

	testCases := []struct {
		ScenarioName  string
		Data          string
		ExpectedData  int
		ExpectedError error
	}{
		{
			ScenarioName:  "Count when 'China'",
			Data:          "China",
			ExpectedData:  178,
			ExpectedError: nil,
		},
		{
			ScenarioName:  "Count when 'Brazil'",
			Data:          "Brazil",
			ExpectedData:  45,
			ExpectedError: nil,
		},
		{
			ScenarioName:  "Count when 'France'",
			Data:          "France",
			ExpectedData:  37,
			ExpectedError: nil,
		},
		{
			ScenarioName:  "Count when Unknown",
			Data:          "Unknown",
			ExpectedData:  0,
			ExpectedError: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.ScenarioName, func(t *testing.T) {
			output := tickets.GetTotalTickets(testCase.Data)
			require.Equal(t, testCase.ExpectedData, output)
		})
	}
}

func TestGetCountByPeriod(t *testing.T) {
	FillTicketList()

	testCases := []struct {
		ScenarioName  string
		Data          string
		ExpectedData  int
		ExpectedError error
	}{
		{
			ScenarioName:  "Count when noite",
			Data:          "21:30",
			ExpectedData:  151,
			ExpectedError: nil,
		},
		{
			ScenarioName:  "Count when tarde",
			Data:          "14:30",
			ExpectedData:  289,
			ExpectedError: nil,
		},
		{
			ScenarioName:  "Count when manha",
			Data:          "11:15",
			ExpectedData:  256,
			ExpectedError: nil,
		},
		{
			ScenarioName:  "Count when inicio manha",
			Data:          "4:25",
			ExpectedData:  304,
			ExpectedError: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.ScenarioName, func(t *testing.T) {
			output, err := tickets.GetCountByPeriod(testCase.Data)
			require.Nil(t, err)
			require.Equal(t, testCase.ExpectedData, output)
		})
	}
	t.Run("Count when all", func(t *testing.T) {

		sumInicioManha, errInicioManha := tickets.GetCountByPeriod("3:30")
		sumManha, errManha := tickets.GetCountByPeriod("10:17")
		sumTarde, errTarde := tickets.GetCountByPeriod("17:18")
		sumNoite, errNoite := tickets.GetCountByPeriod("21:09")

		require.Nil(t, errInicioManha, errManha, errTarde, errNoite)

		sum := sumInicioManha + sumManha + sumTarde + sumNoite
		require.Equal(t, sum, 1000)
	})
}

func TestAverageDestination(t *testing.T) {
	FillTicketList()

	testCases := []struct {
		ScenarioName  string
		Data          string
		ExpectedData  float64
		ExpectedError error
	}{
		{
			ScenarioName:  "Count when 'Mongolia'",
			Data:          "Mongolia",
			ExpectedData:  0.3,
			ExpectedError: nil,
		},
		{
			ScenarioName:  "Count when 'Brazil'",
			Data:          "Brazil",
			ExpectedData:  4.5,
			ExpectedError: nil,
		},
		{
			ScenarioName:  "Count when 'Poland'",
			Data:          "Poland",
			ExpectedData:  3.4000000000000004,
			ExpectedError: nil,
		},
		{
			ScenarioName:  "Count when 'Japan'",
			Data:          "Japan",
			ExpectedData:  1.4000000000000001,
			ExpectedError: nil,
		},
		{
			ScenarioName:  "Count when Unknown",
			Data:          "Unknown",
			ExpectedData:  float64(0),
			ExpectedError: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.ScenarioName, func(t *testing.T) {
			output, err := tickets.AverageDestination(testCase.Data)
			require.Nil(t, err)
			require.Equal(t, testCase.ExpectedData, output)
		})
	}
}
