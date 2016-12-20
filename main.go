package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prinsmike/httprouter"
)

var httpPort string = os.Getenv("APP_HTTP_PORT")

var routes httprouter.Routes

func init() {
	if httpPort == "" {
		httpPort = "1337"
	}
}

func main() {
	r := httprouter.New(routes)

	startupLog()

	err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), r)
	if err != nil {
		log.Println("Fatal error. Could not start the service.")
		log.Fatalln(err)
	}
}

func startupLog() {
	log.Println("Starting the service.")
	fmt.Printf("HTTP Port: %s\n", httpPort)
}

func accessLog(r *http.Request) {
	log.Printf("Access: %s %s\n", r.URL.Path, r.Method)
}
