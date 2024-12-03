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

func MontaPeriodoTicket(time string) (int, error) {
	parts := strings.Split(time, ":")
	if len(parts) != 2 {
		return 0, errors.New(fmt.Sprint("Linha com formato inválido:", time))
	}

	horaStr := strings.TrimSpace(parts[0])

	hora, err := strconv.Atoi(horaStr)
	if err != nil {
		return 0, errors.New(fmt.Sprint("Erro ao converter hora:", err))
	}

	return hora, nil
}

func GetCountTicketsPeriodo(horaMin, horaMax int) (int, error) {

	var countTickets int

	for _, ticket := range listTickets {
		hora, err := MontaPeriodoTicket(ticket.horario)
		if err != nil {
			return 0, err
		}

		if hora >= horaMin && hora < horaMax {
			countTickets++
		}
	}

	return countTickets, nil
}

const (
	inicioManhaComeca = 0
	manhaComeca       = 7
	tardeComeca       = 13
	noiteComeca       = 20
	noiteAcaba        = 25
)

func GetCountByPeriod(time string) (int, error) {
	hora, err := MontaPeriodoTicket(time)
	if err != nil {
		return 0, err
	}

	switch {
	case hora >= inicioManhaComeca && hora < manhaComeca:
		return GetCountTicketsPeriodo(inicioManhaComeca, manhaComeca)
	case hora >= manhaComeca && hora < tardeComeca:
		return GetCountTicketsPeriodo(manhaComeca, tardeComeca)
	case hora >= tardeComeca && hora < noiteComeca:
		return GetCountTicketsPeriodo(tardeComeca, noiteComeca)
	case hora >= noiteComeca && hora < noiteAcaba:
		return GetCountTicketsPeriodo(noiteComeca, noiteAcaba)
	}

	return 0, errors.New("O Periodo nao é valido!")
}

func AverageDestination(destination string) (float64, error) {

	var listPais []Ticket
	for _, itemTicket := range listTickets {
		if itemTicket.paisDestino == destination {
			listPais = append(listPais, itemTicket)
		}
	}

	percentage := (float64(len(listPais)) / float64(len(listTickets))) * 100

	return percentage, nil
}
