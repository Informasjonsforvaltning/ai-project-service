package service

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/Informasjonsforvaltning/ai-project-service/model"
)

// CSVData represents a row of data in the CSV file
type CSVData []string

// CSVService defines the interface for reading data from a CSV file
type CSVService interface {
	ReadCSVData() ([][]string, error)
	MapCsvResponse(data [][]string) []model.AIProject
}

// CSVServiceImpl is an implementation of the CSVService interface
type CSVServiceImpl struct {
}

func InitService() *CSVServiceImpl {
	service := CSVServiceImpl{}
	return &service
}

// ReadCSVData reads data from a CSV file and returns the data as a slice of CSVData
func (service *CSVServiceImpl) ReadCSVData(filepath string) ([]model.AIProject, error) {
	// Open the CSV file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the data from the file into a CSV reader
	reader := csv.NewReader(file)

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
		// Append the row to the data and include headers to the row
		data = append(data, row)
	}

	return service.MapCsvResponse(data), nil
}

func (service *CSVServiceImpl) MapCsvResponse(data [][]string) []model.AIProject {
	var projects []model.AIProject
	for _, row := range data[1:] {
		projects = append(projects, model.AIProject{
			ID:             row[0],
			Prosjekteier:   row[1],
			Prosjekttittel: row[2],
			Departement:    row[3],
			Eiertype:       row[4],
			Kontaktperson:  row[5],
			//Sektor:                  row[6], IKKE I BRUK
			ProsjektBeskrivelse:       row[7],
			ProsjektFormål:            row[8],
			Prosjektstart:             row[9],
			Prosjektslutt:             row[10],
			TilknyttedeOrganisasjoner: row[11],
			InnleideKonsulenter:       row[12],
			LenkeTilProsjekt:          row[13],
			Status:                    row[14],
			TypeData:                  row[15],
			Datakilde:                 row[16],
			Modellutvikling:           row[17],
			Klassifisering:            row[18],
		})

	}
	return projects
}
