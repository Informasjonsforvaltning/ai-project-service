package tests

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/Informasjonsforvaltning/ai-project-service/service"
)

func TestReadCSVData(t *testing.T) {
	// Create a mock CSV file
	file, err := os.Create("mock.csv")
	if err != nil {
		t.Errorf("Error creating mock CSV file: %v", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)

	csvData := [][]string{
		{"ID", "Prosjekteier", "Prosjekttittel", "Departement", "Eiertype", "Sektor", "ProsjektBeskrivelse", "ProsjektFormål", "Prosjektstart", "Prosjektslutt", "TilknyttedeOrganisasjoner", "InnleideKonsulenter", "LenkeTilProsjekt", "Status", "TypeData", "Datakilde", "Modellutvikling", "Klassifisering"},
		{"1", "Prosjekteier 1", "Prosjekttittel 1", "Departement 1", "Eiertype 1", "Sektor 1", "ProsjektBeskrivelse 1", "ProsjektFormål 1", "Prosjektstart 1", "Prosjektslutt 1", "TilknyttedeOrganisasjoner 1", "InnleideKonsulenter 1", "LenkeTilProsjekt 1", "Status 1", "TypeData 1", "Datakilde 1", "Modellutvikling 1", "Klassifisering 1"},
		{"2", "Prosjekteier 2", "Prosjekttittel 2", "Departement 2", "Eiertype 2", "Sektor 2", "ProsjektBeskrivelse 2", "ProsjektFormål 2", "Prosjektstart 2", "Prosjektslutt 2", "TilknyttedeOrganisasjoner 2", "InnleideKonsulenter 2", "LenkeTilProsjekt 2", "Status 2", "TypeData 2", "Datakilde 2", "Modellutvikling 2", "Klassifisering 2"},
	}
	for _, record := range csvData {
		if err := w.Write(record); err != nil {
			t.Errorf("Error writing mock CSV data: %v", err)
		}
	}
	w.Flush()

	// Create a new CSVServiceImpl
	csvService := service.CSVServiceImpl{}

	// Call the ReadCSVData method with the mock data
	filepath := "mock.csv"
	projects, err := csvService.ReadCSVData(filepath)
	if err != nil {
		t.Errorf("Error reading CSV data: %v", err)
	}

	// Assert that the number of projects returned is equal to the number of rows in the mock data
	if len(projects) != len(csvData)-1 {
		t.Errorf("Expected %d projects, but got %d", len(csvData)-1, len(projects))
	}
}

func TestMapCsvResponse(t *testing.T) {
	service := &service.CSVServiceImpl{}

	data := [][]string{
		{"ID", "Prosjekteier", "Prosjekttittel", "Departement", "Eiertype", "Sektor", "ProsjektBeskrivelse", "ProsjektFormål", "Prosjektstart", "Prosjektslutt", "TilknyttedeOrganisasjoner", "InnleideKonsulenter", "LenkeTilProsjekt", "Status", "TypeData", "Datakilde", "Modellutvikling", "Klassifisering"},
		{"1", "owner1", "title1", "department1", "ownertype1", "sector1", "description1", "purpose1", "start1", "end1", "org1", "consultant1", "link1", "status1", "data1", "source1", "development1", "classification1"},
		{"2", "owner2", "title2", "department2", "ownertype2", "sector2", "description2", "purpose2", "start2", "end2", "org2", "consultant2", "link2", "status2", "data2", "source2", "development2", "classification2"},
	}

	projects := service.MapCsvResponse(data)

	if len(projects) != 2 {
		t.Fatalf("MapCsvResponse did not return the expected number of projects. Expected: 2, Actual: %d", len(projects))
	}

	if projects[0].ID != "1" {
		t.Errorf("MapCsvResponse did not map the ID correctly. Expected: 1, Actual: %s", projects[0].ID)
	}

	if projects[1].Prosjekteier != "owner2" {
		t.Errorf("MapCsvResponse did not map the Prosjekteier correctly. Expected: owner2, Actual: %s", projects[1].Prosjekteier)
	}
}
