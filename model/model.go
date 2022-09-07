package model

import (
	"encoding/json"
	"log"
)

type Path struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type Trip struct {
	// Origin      string `json:"origin"`
	// Destination string `json:"destination"
	Route     Path   `json:"route"`
	Agency    string `json:"agency"`
	Supplier  string `json:"supplier"`
	Airline   string `json:"airline"`
	BasePrice int    `json:"base_price"`
	MarkUp    int    `json:"markup"`
}

func TripFromJson(jsonData []byte) Trip {
	var trip Trip
	err := json.Unmarshal(jsonData, &trip)
	if err != nil {
		log.Fatalf("Error occurred during decoding. Error: %s", err.Error())
	}
	return trip
}

func (self *Trip) ToString() string {
	json, err := json.Marshal(self)
	if err != nil {
		log.Fatalf("Error occurred during encoding. Error: %s", err.Error())
	}
	return string(json)
}

func (self *Trip) PrettyString() string {
	b, err := json.MarshalIndent(self, "", "   ")
	if err != nil {
		log.Fatalf("Error occurred during encoding. Error: %s", err.Error())
	}
	return string(b)
}
