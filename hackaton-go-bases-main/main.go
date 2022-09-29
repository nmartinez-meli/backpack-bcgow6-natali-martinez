package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []service.Ticket
	var ticket service.Ticket
	archivo := file.File{Path: "tickets.csv"}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	tickets, err := archivo.Read()
	if err != nil {
		panic(err)
	}
	b := service.NewBookings(tickets)

	ticket, err = b.Read(24)
	if err != nil {
		panic(err)
	}
	fmt.Println(ticket)
}
