package file

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type File struct {
	Path string
}

func (f *File) Read() ([][]string, error) {
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

	return records, nil
}

func (f *File) Write(ticketFiels []string) error {

	openFile, err := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer openFile.Close()
	if err != nil {
		return err
	}
	openFile.WriteString("\n" + strings.Join(ticketFiels, ","))
	if err != nil {
		return err
	}
	return nil
}
func (f *File) WriteAll(tickets [][]string) error {
	os.Remove(f.Path)
	openFile, err := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer openFile.Close()
	if err != nil {
		return err
	}
	w := csv.NewWriter(openFile)
	defer w.Flush()

	for _, record := range tickets {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	return nil
}
