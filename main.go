package main

import (
	"fmt"
	data "snapp-trip/data"
	model "snapp-trip/model"
)

const (
	path1 = "data/sample.json"
	path2 = "data/sample2.json"
	link2 = "https://filebin.net/x3ky3d43jerain9a"
)

func main() {
	json_data := data.GetFromPath(path2)
	flight := model.TripFromJson(json_data)
	// println(flight.Route)
	fmt.Println(flight.Agency)
	fmt.Println(flight.Route)
	println(flight.ToString())
	// fmt.Println(flight.PrettyString())
	body := data.GetFromLink(link2)
	println(string(body))
	// flight_link := model.TripFromJson(body)
	// println(flight_link.ToString())
	// add_file()
}

// func add_file() {
// 	http.HandleFunc("/", greet)
// 	http.ListenAndServe(":8080", nil)
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	file_data, err := ioutil.ReadFile(path2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// return string(file_data)
// 	fmt.
// }
