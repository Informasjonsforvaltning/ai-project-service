package model

type AIProject struct {
	ID                        string `json:"id"`
	Prosjekteier              string `json:"prosjekteier"`
	Prosjekttittel            string `json:"prosjekttittel"`
	Departement               string `json:"departement"`
	Eiertype                  string `json:"eiertype"`
	Kontaktperson             string `json:"kontaktperson"`
	Sektor                    string `json:"sektor"`
	ProsjektBeskrivelse       string `json:"prosjektBeskrivelse"`
	ProsjektFormål            string `json:"prosjektFormaal"`
	Prosjektstart             string `json:"prosjektstart"`
	Prosjektslutt             string `json:"prosjektslutt"`
	TilknyttedeOrganisasjoner string `json:"tilknyttedeOrganisasjoner"`
	InnleideKonsulenter       string `json:"innleideKonsulenter"`
	LenkeTilProsjekt          string `json:"lenkeTilProsjekt"`
	Status                    string `json:"status"`
	TypeData                  string `json:"typeData"`
	Datakilde                 string `json:"datakilde"`
	Modellutvikling           string `json:"modellutvikling"`
	Klassifisering            string `json:"klassifisering"`
}
