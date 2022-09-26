package main

import (
	"fmt"
	model "snapp-trip/model"
	server "snapp-trip/server"
)

const (
	path1 = "data/samples/sample.json"
	path2 = "data/samples/create_rule_request.json"
	link2 = "https://filebin.net/x3ky3d43jerain9a"
)

func main() {
	json_data := server.GetFromPath(path2)
	flight := model.TripsFromJson(json_data)
	fmt.Println(len(flight))
	fmt.Println(flight[0].Agency)
	fmt.Println(flight[0].Origin)
	println(flight[0].ToString())
	fmt.Println(flight[0].PrettyString())

	server.Connect()

	// flight[0]
	// body := data.GetFromLink(link2)
	// println(string(body))
	// flight_link := model.TripFromJson(body)
	// println(flight_link.ToString())
	// add_file()
}
