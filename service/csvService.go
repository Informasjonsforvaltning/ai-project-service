package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// CSVHeaders defines the headers in the CSV file
var CSVHeaders = []string{
	"ID",
	"Prosjekteier",
	"Prosjekttittel",
	"Departement",
	"Eiertype",
	"Kontaktperson",
	"Sektor (intern | ikke publisering)",
	"Beskrivelse av prosjekt",
	"Formål med prosjekt",
	"Prosjektstart",
	"Prosjektslutt",
	"Tilknyttede organisasjoner",
	"Bruk av innleide konsulenter",
	"Lenke til prosjekt",
	"Status",
	"Type data",
	"Datakilde",
	"Modellutvikling",
	"Klassifisering",
}

// CSVData represents a row of data in the CSV file
type CSVData []string

// CSVService defines the interface for reading data from a CSV file
type CSVService interface {
	ReadCSVData() ([][]string, error)
}

// CSVServiceImpl is an implementation of the CSVService interface
type CSVServiceImpl struct {
}

func InitService() *CSVServiceImpl {
	service := CSVServiceImpl{}
	return &service
}

// ReadCSVData reads data from a CSV file and returns the data as a slice of CSVData
func (service *CSVServiceImpl) ReadCSVData() ([][]string, error) {
	// Open the CSV file
	file, err := os.Open("ai_projects_norwegian_state - Oversatt_v1.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the data from the file into a CSV reader
	reader := csv.NewReader(file)

	// Read the header row
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Verify that the header row matches the expected headers
	if len(header) != len(CSVHeaders) {
		return nil, fmt.Errorf("incorrect number of headers in CSV file")
	}
	for i := range header {
		if header[i] != CSVHeaders[i] {
			return nil, fmt.Errorf("incorrect header in CSV file: expected %s, got %s", CSVHeaders[i], header[i])
		}
	}

	// Read the rest of the data from the file
	var data [][]string
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}

	return data, nil
}
