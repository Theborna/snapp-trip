package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	// defer db.Close()
	connection = db

	// err = db.Ping()
	CheckError(err)

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
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
