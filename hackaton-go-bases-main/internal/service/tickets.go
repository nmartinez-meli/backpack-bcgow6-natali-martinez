package service

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
	file    file.File
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(path string) (Bookings, error) {
	archivo := file.File{Path: path}
	records, err := archivo.Read()
	if err != nil {
		return nil, err
	}
	tickets := []Ticket{}
	for _, ticket := range records {
		id, err := strconv.Atoi(ticket[0])
		if err != nil {
			return &bookings{}, err
		}
		price, err := strconv.Atoi(ticket[5])
		if err != nil {
			return &bookings{}, err
		}
		tickets = append(tickets, Ticket{
			Id:          id,
			Names:       ticket[1],
			Email:       ticket[2],
			Destination: ticket[3],
			Date:        ticket[4],
			Price:       price,
		})
	}
	return &bookings{Tickets: tickets, file: archivo}, nil
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	valuesTicket := reflect.ValueOf(t)

	ticketFiels := []string{}
	for i := 0; i < valuesTicket.NumField(); i++ {
		ticketFiels = append(ticketFiels, fmt.Sprintf("%v", valuesTicket.Field(i).Interface()))
	}
	err := b.file.Write(ticketFiels)
	if err != nil {
		return Ticket{}, err
	}
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	sort.Slice(b.Tickets, func(i, j int) bool {
		return b.Tickets[i].Id < b.Tickets[j].Id
	})
	i := sort.Search(len(b.Tickets), func(i int) bool { return b.Tickets[i].Id >= id && i < len(b.Tickets)-1 })
	if b.Tickets[i].Id != id {
		return Ticket{}, errors.New("error: el ticket no se encuentra registrado")
	}
	return b.Tickets[i], nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	var newBookings [][]string
	for i, record := range b.Tickets {
		var row []string
		if record.Id == id {
			row = []string{fmt.Sprint(record.Id), t.Names, t.Email, t.Destination, t.Date, fmt.Sprint(t.Price)}
			b.Tickets[i] = t
		} else {
			row = []string{fmt.Sprint(record.Id), record.Names, record.Email, record.Destination, record.Date, fmt.Sprint(record.Price)}
		}
		newBookings = append(newBookings, row)
	}

	err := b.file.WriteAll(newBookings)
	if err != nil {
		return Ticket{}, err
	}
	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	var newBookings [][]string
	var tickets []Ticket

	for _, record := range b.Tickets {
		var row []string
		if record.Id != id {
			row = []string{fmt.Sprint(record.Id), record.Names, record.Email, record.Destination, record.Date, fmt.Sprint(record.Price)}
			tickets = append(tickets, record)
			newBookings = append(newBookings, row)
		}
	}

	err := b.file.WriteAll(newBookings)
	if err != nil {
		return id, err
	}
	return id, nil
}
