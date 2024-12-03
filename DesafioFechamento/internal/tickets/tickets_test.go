package tickets_test

import (
	"testing"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	tickets.PreencherListTickets()

	t.Run("Quantidade certa de tickets", func(t *testing.T) {
		total := tickets.GetTotalTickets("Brazil")
		expected := 45
		require.Equal(t, total, expected)
	})
	t.Run("Quantidade errada de tickets", func(t *testing.T) {
		total := tickets.GetTotalTickets("Brazil")
		expected := 44
		require.NotEqual(t, total, expected)
	})

	t.Run("Quantidade certa do periodo", func(t *testing.T) {
		countHorario, err := tickets.GetCountByPeriod("21:30")
		expected := 151
		require.Nil(t, err)
		require.Equal(t, countHorario, expected)
	})

	t.Run("Quantidade errada do periodo", func(t *testing.T) {
		countHorario, err := tickets.GetCountByPeriod("14:30")
		expected := 190
		require.Nil(t, err)
		require.NotEqual(t, countHorario, expected)
	})

	t.Run("Media certa do brazil", func(t *testing.T) {
		porcentagem, err := tickets.AverageDestination("Brazil")
		expected := 4.5
		require.Nil(t, err)
		require.Equal(t, porcentagem, expected)
	})

	t.Run("Media errada do brazil", func(t *testing.T) {
		porcentagem, err := tickets.AverageDestination("Brazil")
		expected := 7.2
		require.Nil(t, err)
		require.NotEqual(t, porcentagem, expected)
	})
}
