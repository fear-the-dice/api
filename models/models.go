package models

import (
	"os"
)

var (
	uri string
)

func Setup() {
	uri := os.Getenv("MONGOLAB_URI")
	if uri == "" {
		uri = "mongodb://localhost/heroku_app37083199"
	}
}
