package model

import (
	"encoding/json"
	"log"
)

func TripsFromJson(jsonData []byte) []Trip {
	var trips []Trip
	err := json.Unmarshal(jsonData, &trips)
	if err != nil {
		log.Fatalf("Error occurred during decoding. Error: %s", err.Error())
	}
	return trips
}

func RulesFromJson(jsonData []byte) []Rule {
	var rules []Rule
	err := json.Unmarshal(jsonData, &rules)
	if err != nil {
		log.Fatalf("Error occurred during decoding. Error: %s", err.Error())
	}
	return rules
}

func ToString(__self interface{}) []byte {
	json, err := json.Marshal(__self)
	if err != nil {
		log.Fatalf("Error occurred during encoding. Error: %s", err.Error())
	}
	return json
}

func PrettyString(__self interface{}) string {
	b, err := json.MarshalIndent(__self, "", "   ")
	if err != nil {
		log.Fatalf("Error occurred during encoding. Error: %s", err.Error())
	}
	return string(b)
}

func (__self *Trip) ToString() []byte {
	return ToString(__self)
}

func (__self *Trip) PrettyString() string {
	return PrettyString(__self)
}

func (__self *Rule) ToString() []byte {
	return ToString(__self)
}

func (__self *Rule) PrettyString() string {
	return PrettyString(__self)
}

func (__self *Rule) Info() {
	log.Printf("the Routes are %v with  type %T ", __self.Routes, __self.Routes)
	log.Printf("the Airlines are %v with  type %T ", __self.Airlines, __self.Airlines)
	log.Printf("the Agencies are %v with  type %T ", __self.Agencies, __self.Agencies)
	log.Printf("the Suppliers are %v with  type %T ", __self.Suppliers, __self.Suppliers)
	log.Printf("the AmountType are %v with  type %T ", __self.AmountType, __self.AmountType)
	log.Printf("the AmountValue are %v with  type %T ", __self.AmountValue, __self.AmountValue)
}

func (__self *Response) ToString() []byte {
	return ToString(__self)
}

func (__self *Response) PrettyString() string {
	return PrettyString(__self)
}
