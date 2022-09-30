package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var ticket, newTicket service.Ticket
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	b, err := service.NewBookings("tickets.csv")
	if err != nil {
		panic(err)
	}

	ticket, err = b.Read(24)
	if err != nil {
		println(err)
	}
	fmt.Println(ticket)
	t := ticket
	t.Id = 1001
	newTicket, err = b.Create(t)
	if err != nil {
		println(err)
	}
	fmt.Printf("el ticket %v fue creado con exito\n", newTicket)
	t = ticket
	t.Names = "NATALI"
	newTicket, err = b.Update(24, t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("el ticket %v fue actualizado con exito\n", newTicket)
	ticketIndex, err := b.Delete(25)
	if err != nil {
		fmt.Printf("el ticket %d no pudo ser eliminado\n error: %v\n", ticketIndex, err)
	}
	fmt.Printf("el ticket con id (%d) fue eliminado con exito\n", ticketIndex)
}
