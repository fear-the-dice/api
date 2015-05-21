package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/fear-the-dice/api/controllers"
)

func main() {
	var (
		host = flag.String("HOST", "keepiteasy.net", "The host name the server is running on")
		port = flag.String("PORT", ":3000", "The server port")
	)

	flag.Parse()

	config := controllers.ControllerConfig{
		Host: *host,
		Port: *port,
	}

	controller := controllers.NewController(&config)
	mux := http.NewServeMux()

	controller.Attach(mux)

	if err := http.ListenAndServe(*port, mux); err != nil {
		log.Fatal(err)
	}
}
