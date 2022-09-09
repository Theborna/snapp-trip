package main

import (
	"fmt"
	data "snapp-trip/data"
	model "snapp-trip/model"
)

const (
	path1 = "data/sample.json"
	path2 = "data/sample2.json"
	path3 = "data/create_rule_request.json"
	link2 = "https://filebin.net/x3ky3d43jerain9a"
)

func main() {
	json_data := data.GetFromPath(path1)
	flight := model.TripsFromJson(json_data)
	fmt.Println(flight[0].Agency)
	fmt.Println(flight[0].Origin)
	println(flight[0].ToString())
	fmt.Println(flight[0].PrettyString())

	rule_json_data := data.GetFromPath(path3)
	rules := model.RulesFromJson(rule_json_data)
	fmt.Println(rules[0].Airlines)
	fmt.Println(rules[0].Routes)
	println(rules[0].ToString())
	fmt.Println(rules[0].PrettyString())
	rules[0].Info()
	// flight[0]
	// body := data.GetFromLink(link2)
	// println(string(body))
	// flight_link := model.TripFromJson(body)
	// println(flight_link.ToString())
	// add_file()
}
