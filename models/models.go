package models

import (
	"os"
)

var (
	uri string
)

// User stores a users name
type (
	Base struct {
		ID         string `bson:"id" json:"id"`
		Initiative int    `bson:"initiative"`
		AC         int    `bson:"ac"`
		HP         int    `bson:"hp"`
		Health     int    `bson:"health"`
		Damage     int    `bson:"damage"`
		Speed      string `bson:"speed"`
	}
)

func Setup() {
	uri := os.Getenv("MONGOLAB_URI")
	if uri == "" {
		uri = "mongodb://localhost/heroku_app37083199"
	}
}
