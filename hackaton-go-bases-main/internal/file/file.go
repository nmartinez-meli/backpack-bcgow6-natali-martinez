package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	openFile, err := os.Open(f.Path)

	if err != nil {

		return nil, err
	}
	defer openFile.Close()

	r := csv.NewReader(openFile)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	tickets := []service.Ticket{}
	for _, ticket := range records {
		id, err := strconv.Atoi(ticket[0])
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, service.Ticket{
			Id:          id,
			Names:       ticket[1],
			Email:       ticket[2],
			Destination: ticket[3],
			Date:        ticket[4],
		})

	}
	return tickets, nil
}

func (f *File) Write(ticket service.Ticket) error {
	openFile, err := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(openFile)

	valuesTicket := reflect.ValueOf(ticket)

	ticketFiels := []string{}
	for i := 0; i < valuesTicket.NumField(); i++ {
		ticketFiels = append(ticketFiels, fmt.Sprintf("%v", valuesTicket.Field(i).Interface()))
	}
	if err := writer.Write(ticketFiels); err != nil {
		return err
	}
	return nil
}
