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
		fmt.Println(err)
	} else {

		fmt.Println(ticket)
	}
	t := service.Ticket{Id: 1001, Names: "Natali Martinez", Email: "natali.martinez@mercadolibre.com", Destination: "fortaleza", Date: "7-12-2022", Price: 4500}

	newTicket, err = b.Create(t)
	if err != nil {

		fmt.Println(err)
	} else {
		fmt.Printf("el ticket %v fue creado con exito\n", newTicket)
	}
	t = ticket
	t.Names = "NATALI"
	newTicket, err = b.Update(24, t)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("el ticket %v fue actualizado con exito\n", newTicket)
	}
	ticketIndex, err := b.Delete(25)
	if err != nil {
		fmt.Printf("el ticket %d no pudo ser eliminado\n error: %v\n", ticketIndex, err)
	} else {

		fmt.Printf("el ticket con id (%d) fue eliminado con exito\n", ticketIndex)
	}
}
