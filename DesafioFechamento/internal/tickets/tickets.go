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
)

const (
	morningStart     = 0
	morningStartTime = 7
	afternoonStart   = 13
	eveningStart     = 20
	eveningEnd       = 25
)

const (
	fieldID = iota
	fieldName
	fieldEmail
	fieldDestination
	fieldTime
	fieldPrice
)

type Ticket struct {
	id          int
	name        string
	email       string
	destination string
	time        string
	price       float32
}

var ticketList []Ticket

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

func init() {
	FillTicketList()
}

func BuildTicket(line string) (Ticket, error) {
	parts := strings.Split(line, ",")
	if len(parts) != PartsTicket {
		var formatError ErrorInvalidLineFormat
		formatError.message = fmt.Sprintf("Error: Invalid line format: %s ", line)

		return Ticket{}, formatError
	}

	idStr := strings.TrimSpace(parts[fieldID])
	nameTicket := strings.TrimSpace(parts[fieldName])
	emailTicket := strings.TrimSpace(parts[fieldEmail])
	destinationTicket := strings.TrimSpace(parts[fieldDestination])
	timeTicket := strings.TrimSpace(parts[fieldTime])
	priceStr := strings.TrimSpace(parts[fieldPrice])

	idTicket, err := strconv.Atoi(idStr)
	if err != nil {
		convertError := errors.New("Error converting id:" + err.Error())
		return Ticket{}, convertError
	}

	priceTicket, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		convertError := errors.New("Error converting price: " + err.Error())
		return Ticket{}, convertError
	}

	ticket := Ticket{
		id:          idTicket,
		name:        nameTicket,
		email:       emailTicket,
		destination: destinationTicket,
		time:        timeTicket,
		price:       float32(priceTicket),
	}

	return ticket, nil
}

func validateTicket(ticket Ticket) error {
	if ticket.id == 0 || ticket.name == "" || ticket.email == "" || ticket.destination == "" || ticket.time == "" || ticket.price == 0.0 {
		return errors.New("Ticket with empty information - name:" + ticket.name)
	}

	return nil
}

func AddTicket(ticket Ticket) error {
	for _, itemTicket := range ticketList {
		if itemTicket.id == ticket.id {
			var errorAlreadyExists ErrorInvalidLineFormat
			errorAlreadyExists.message = fmt.Sprintf("Error: Ticket already exists - id: %d ", itemTicket.id)

			return errorAlreadyExists
		}
	}

	err := validateTicket(ticket)
	if err != nil {
		return err
	}

	ticketList = append(ticketList, ticket)
	return nil
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

		errAdd := AddTicket(ticket)
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
	var countCountry int
	for _, itemTicket := range ticketList {
		if itemTicket.destination == destination {
			countCountry++
		}
	}

	percentage := (float64(countCountry) / float64(len(ticketList))) * 100
	return percentage, nil
}
