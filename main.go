package main

import (
	"flag"
	"log"
	"net/http"

	"snapp-trip/api"

	"time"
)

var (
	httpServerAddr              = flag.String("localhost", ":5500", "TCP address that the http server will listen on")
	httpServerReadTimeout       = flag.Duration("http-server.read-timeout", 5*time.Minute, "duration before timing out when reading request, includes body")
	httpServerReadHeaderTimeout = flag.Duration("http-server.read-header-timeout", 30*time.Second, "duration before timing out when reading headers")
	httpServerWriteTimeout      = flag.Duration("http-server.write-timeout", 30*time.Second, "duration before timing out when writing response")
)

func main() {
	server := &http.Server{
		Addr:              *httpServerAddr,
		ReadTimeout:       *httpServerReadTimeout,
		ReadHeaderTimeout: *httpServerReadHeaderTimeout,
		WriteTimeout:      *httpServerWriteTimeout,
		Handler:           api.NewServer(),
	}
	// database.Connect()
	// city := model.City("asdad")
	// log.Printf("got %s", city.Validate())
	log.Printf("Starting server on %s\n", *httpServerAddr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
