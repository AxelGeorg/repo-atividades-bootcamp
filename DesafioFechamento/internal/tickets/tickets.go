package tickets

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ErrorLinhaFormatoInvalido struct {
	message string
}

func (e ErrorLinhaFormatoInvalido) Error() string {
	return e.message
}

type ErrorTicketJaExisteLista struct {
	message string
}

func (e ErrorTicketJaExisteLista) Error() string {
	return e.message
}

type Ticket struct {
	id          int
	nome        string
	email       string
	paisDestino string
	horario     string
	preco       float32
}

var listTickets []Ticket

func init() {
	file, err := os.Open("tickets.csv")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	defer file.Close()

	PreencherListTickets(file)
}

func MontaTicket(linha string) (*Ticket, error) {
	parts := strings.Split(linha, ",")
	if len(parts) != 6 {
		var errorFormato ErrorLinhaFormatoInvalido
		errorFormato.message = fmt.Sprintf("Erro: Linha com formato inválido: %s ", linha)

		return nil, errorFormato
	}

	idStr := strings.TrimSpace(parts[0])
	nomeTicket := strings.TrimSpace(parts[1])
	emailTicket := strings.TrimSpace(parts[2])
	paisTicket := strings.TrimSpace(parts[3])
	horarioTicket := strings.TrimSpace(parts[4])
	precoStr := strings.TrimSpace(parts[5])

	idTicket, err := strconv.Atoi(idStr)
	if err != nil {
		errorConvert := errors.New(fmt.Sprintf("Erro ao converter id:", err))
		return nil, errorConvert
	}

	precoTicket, err := strconv.ParseFloat(precoStr, 32)
	if err != nil {
		errorConvert := errors.New(fmt.Sprintf("Erro ao converter preco:", err))
		return nil, errorConvert
	}

	ticket := Ticket{
		id:          idTicket,
		nome:        nomeTicket,
		email:       emailTicket,
		paisDestino: paisTicket,
		horario:     horarioTicket,
		preco:       float32(precoTicket),
	}

	return &ticket, nil
}

func AdicionaTicket(ticket Ticket) (bool, error) {
	for _, itemTicket := range listTickets {
		if itemTicket.id == ticket.id {
			var errorJaExiste ErrorLinhaFormatoInvalido
			errorJaExiste.message = fmt.Sprintf("Erro: Ticket already exists - id: %d ", itemTicket.id)

			return false, errorJaExiste
		}
	}

	if ticket.id == 0 || ticket.nome == "" || ticket.email == "" || ticket.paisDestino == "" || ticket.horario == "" || ticket.preco == 0.0 {
		return false, errors.New(fmt.Sprintf("Ticket com informacao zerada: ", ticket))
	}

	listTickets = append(listTickets, ticket)

	return true, nil
}

func PreencherListTickets(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ticket, err := MontaTicket(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		_, errAdd := AdicionaTicket(*ticket)
		if errAdd != nil {
			fmt.Println(errAdd.Error())
			continue
		}
	}
}

func GetTotalTickets(destination string) int {

	var countTickets int

	for _, ticket := range listTickets {
		if ticket.paisDestino == destination {
			countTickets++
		}
	}

	return countTickets
}

const (
	inicioManhaMin = 0
	inicioManhaMax = 6
	manhaMin       = 7
	manhaMax       = 12
	tardeMin       = 13
	tardeMax       = 19
	noiteMin       = 20
	noiteMax       = 23
)

func MontaPeriodoTicket(time string) (int, int, error) {
	parts := strings.Split(time, ":")
	if len(parts) != 2 {
		return 0, 0, errors.New(fmt.Sprint("Linha com formato inválido:", time))
	}

	horaStr := strings.TrimSpace(parts[0])
	minStr := strings.TrimSpace(parts[1])

	hora, err := strconv.Atoi(horaStr)
	if err != nil {
		return 0, 0, errors.New(fmt.Sprint("Erro ao converter hora:", err))
	}

	min, err := strconv.Atoi(minStr)
	if err != nil {
		return 0, 0, errors.New(fmt.Sprint("Erro ao converter min:", err))
	}

	return hora, min, nil
}

func GetCountByPeriod(time string) (int, error) {
	hora, min, err := MontaPeriodoTicket(time)
	if err != nil {
		return 0, err
	}

	switch {
		case 
	}

	return 0, errors.New("teste")
}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	return 0, errors.New("teste")
}
