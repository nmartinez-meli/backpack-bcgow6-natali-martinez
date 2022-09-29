package service

import (
	"errors"
	"sort"
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
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	return Ticket{}, nil
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
	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
}
