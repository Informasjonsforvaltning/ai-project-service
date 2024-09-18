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
			//Sektor:                  row[5], IKKE I BRUK
			ProsjektBeskrivelse:       row[6],
			ProsjektFormål:            row[7],
			Prosjektstart:             row[8],
			Prosjektslutt:             row[9],
			TilknyttedeOrganisasjoner: row[10],
			InnleideKonsulenter:       row[11],
			LenkeTilProsjekt:          row[12],
			Status:                    row[13],
			TypeData:                  row[14],
			Datakilde:                 row[15],
			Modellutvikling:           row[16],
			Klassifisering:            row[17],
		})

	}
	return projects
}
