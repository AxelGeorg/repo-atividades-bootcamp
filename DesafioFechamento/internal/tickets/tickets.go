package tickets

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PartsPeriod = 2
	PartsTicket = 6

	morningStart     = 0
	morningStartTime = 7
	afternoonStart   = 13
	eveningStart     = 20
	eveningEnd       = 25
)

type ErrorInvalidLineFormat struct {
	message string
}

func (e ErrorInvalidLineFormat) Error() string {
	return e.message
}

type ErrorTicketAlreadyExistsInList struct {
	message string
}

func (e ErrorTicketAlreadyExistsInList) Error() string {
	return e.message
}

type Ticket struct {
	id          int
	name        string
	email       string
	destination string
	time        string
	price       float32
}

var ticketList []Ticket

func BuildTicket(line string) (*Ticket, error) {
	parts := strings.Split(line, ",")
	if len(parts) != PartsTicket {
		var formatError ErrorInvalidLineFormat
		formatError.message = fmt.Sprintf("Error: Invalid line format: %s ", line)

		return nil, formatError
	}

	idStr := strings.TrimSpace(parts[0])
	nameTicket := strings.TrimSpace(parts[1])
	emailTicket := strings.TrimSpace(parts[2])
	destinationTicket := strings.TrimSpace(parts[3])
	timeTicket := strings.TrimSpace(parts[4])
	priceStr := strings.TrimSpace(parts[5])

	idTicket, err := strconv.Atoi(idStr)
	if err != nil {
		convertError := errors.New("Error converting id:" + err.Error())
		return nil, convertError
	}

	priceTicket, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		convertError := errors.New("Error converting price: " + err.Error())
		return nil, convertError
	}

	ticket := Ticket{
		id:          idTicket,
		name:        nameTicket,
		email:       emailTicket,
		destination: destinationTicket,
		time:        timeTicket,
		price:       float32(priceTicket),
	}

	return &ticket, nil
}

func AddTicket(ticket Ticket) (bool, error) {
	for _, itemTicket := range ticketList {
		if itemTicket.id == ticket.id {
			var errorAlreadyExists ErrorInvalidLineFormat
			errorAlreadyExists.message = fmt.Sprintf("Error: Ticket already exists - id: %d ", itemTicket.id)

			return false, errorAlreadyExists
		}
	}

	if ticket.id == 0 || ticket.name == "" || ticket.email == "" || ticket.destination == "" || ticket.time == "" || ticket.price == 0.0 {
		return false, errors.New("Ticket with empty information - name:" + ticket.name)
	}

	ticketList = append(ticketList, ticket)

	return true, nil
}

func GetCountTicketList() int {
	return len(ticketList)
}

func FillTicketList() {
	file, err := os.Open("tickets.csv")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ticket, err := BuildTicket(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		_, errAdd := AddTicket(*ticket)
		if errAdd != nil {
			fmt.Println(errAdd.Error())
			continue
		}
	}
}

func GetTotalTickets(destination string) int {
	var countTickets int

	for _, ticket := range ticketList {
		if ticket.destination == destination {
			countTickets++
		}
	}

	return countTickets
}

func BuildPeriodTicket(time string) (int, error) {
	parts := strings.Split(time, ":")
	if len(parts) != PartsPeriod {
		return 0, errors.New(fmt.Sprint("Line with invalid format:", time))
	}

	hourStr := strings.TrimSpace(parts[0])

	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		return 0, errors.New(fmt.Sprint("Error converting hour:", err))
	}

	return hour, nil
}

func GetCountTicketsPeriod(hourMin, hourMax int) (int, error) {
	var countTickets int

	for _, ticket := range ticketList {
		hour, err := BuildPeriodTicket(ticket.time)
		if err != nil {
			return 0, err
		}

		if hour >= hourMin && hour < hourMax {
			countTickets++
		}
	}

	return countTickets, nil
}

func GetCountByPeriod(time string) (int, error) {
	hour, err := BuildPeriodTicket(time)
	if err != nil {
		return 0, err
	}

	switch {
	case hour >= morningStart && hour < morningStartTime:
		return GetCountTicketsPeriod(morningStart, morningStartTime)
	case hour >= morningStartTime && hour < afternoonStart:
		return GetCountTicketsPeriod(morningStartTime, afternoonStart)
	case hour >= afternoonStart && hour < eveningStart:
		return GetCountTicketsPeriod(afternoonStart, eveningStart)
	case hour >= eveningStart && hour < eveningEnd:
		return GetCountTicketsPeriod(eveningStart, eveningEnd)
	}

	return 0, errors.New("The period is not valid!")
}

func AverageDestination(destination string) (float64, error) {
	var countryList []Ticket
	for _, itemTicket := range ticketList {
		if itemTicket.destination == destination {
			countryList = append(countryList, itemTicket)
		}
	}

	percentage := (float64(len(countryList)) / float64(len(ticketList))) * 100

	return percentage, nil
}
