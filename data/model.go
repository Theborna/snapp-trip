package data

import (
	"snapp-trip/server"
)

type data interface {
	SendToDB() error
}

type Airline struct {
	data
	Name    string `json:"name"`
	NameEng string `json:"name_eng"`
	NameFar string `json:"name_far"`
}

func (airline *Airline) SendToDB() error {
	con := server.GetConnection()
	sqlStatement := `
	INSERT INTO users (name, name_eng, name_fa)
	VALUES ($1, $2, $3)`
	err := con.Raw(sqlStatement, airline.Name, airline.NameEng, airline.NameFar).Error
	return err
}
