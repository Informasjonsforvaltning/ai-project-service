package model

type AIProject struct {
	ID                        string
	Prosjekteier              string
	Prosjekttittel            string
	Departement               string
	Eiertype                  string
	Kontaktperson             string
	Sektor                    string
	ProsjektBeskrivelse       string
	ProsjektFormål            string
	Prosjektstart             string
	Prosjektslutt             string
	TilknyttedeOrganisasjoner string
	InnleideKonsulenter       string
	LenkeTilProsjekt          string
	Status                    string
	TypeData                  string
	Datakilde                 string
	Modellutvikling           string
	Klassifisering            string
}
