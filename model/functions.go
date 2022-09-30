package model

import (
	"encoding/json"
	"fmt"
	"log"
)

func (trip *Trip) SetRule(rules []Rule) {
	defer trip.setPayable()
	max, maxIdx := -1, -1
	for i, v := range rules {
		markup := trip.markupAfterRule(v)
		if max < markup {
			max = markup
			maxIdx = i
		}
	}
	if maxIdx < 0 {
		return
	}
	bestRule := rules[maxIdx]

	trip.RuleId = &bestRule.RuleId
	// setting the price
	if bestRule.AmountType == FIXED {
		trip.MarkUp = bestRule.AmountValue
	} else if bestRule.AmountType == PERCENTAGE {
		trip.MarkUp = trip.BasePrice * bestRule.AmountValue / 100
	}
}

func (trip *Trip) markupAfterRule(rule Rule) int {
	if rule.AmountType == FIXED {
		return rule.AmountValue
	}
	return trip.BasePrice * rule.AmountValue / 100
}

func (trip *Trip) setPayable() {
	trip.PayablePrice = trip.BasePrice + trip.MarkUp
}

func (rule *Rule) ValidForTrip(trip Trip) bool {
	// fmt.Prin tf("trip: %s\n", trip.PrettyString())
	// fmt.Printf("rule: %s\n", rule.PrettyString())
	var agencyValid bool = rule.Agencies == nil
	if rule.Agencies != nil {
		for _, agency := range rule.Agencies {
			if trip.Agency == agency {
				agencyValid = true
				break
			}
		}
	}
	fmt.Printf("agencyValid: %v\n", agencyValid)
	airlineValid := rule.Airlines == nil
	if rule.Airlines != nil {
		for _, airline := range rule.Airlines {
			if trip.Airline == airline {
				airlineValid = true
				break
			}
		}
	}
	fmt.Printf("airlineValid: %v\n", airlineValid)
	supplierValid := rule.Suppliers == nil
	if rule.Suppliers != nil {
		for _, supplier := range rule.Suppliers {
			if trip.Supplier == supplier {
				supplierValid = true
				break
			}
		}
	}
	fmt.Printf("supplierValid: %v\n", supplierValid)
	otherValid := agencyValid && airlineValid && supplierValid
	if !otherValid {
		return false
	}
	var validRoute bool = rule.Routes == nil
	if rule.Routes != nil {
		for _, route := range rule.Routes {
			if (*route.Destination == trip.Destination || route.Destination == nil) && (*route.Origin == trip.Origin || route.Origin == nil) {
				validRoute = true
				break
			}
		}
	}
	fmt.Printf("validOrigin: %v\n", validRoute)
	return otherValid && validRoute
}

// more general functions
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

func (__self *Route) ToString() []byte {
	return ToString(__self)
}

func (__self *Route) PrettyString() string {
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
