package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "bornakh"
	password = "bornakh"
	dbname   = "test"
)

var connection *gorm.DB

func Connect() {
	psqlconn := fmt.Sprintf("host= %s port=%d user = %s password = %s dbname=%s", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	CheckError(err)
	connection = db
}

func GetConnection() *gorm.DB {
	return connection
}

func CheckError(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

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
