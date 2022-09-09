package data

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetFromPath(path string) []byte {
	file_data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	return file_data
}

func GetFromLink(link string) []byte {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("dump", body, 0600)
	return body
}
