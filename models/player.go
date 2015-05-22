package models

import (
	"os"

	"gopkg.in/mgo.v2"
)

// User stores a users name
type (
	Player struct {
		Base
		ID        string `json:"ID"`
		Name      string `json:"name"`
		Character string `json:"character"`
	}
	Players []*Player
)

func NewPlayer() *Player {
	return &Player{}
}

func PopulatePlayers() (*Players, error) {
	uri := os.Getenv("MONGOLAB_URI")

	if uri == "" {
		uri = "mongodb://localhost/heroku_app37083199"
	}

	Session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	defer Session.Close()

	Session.SetSafe(&mgo.Safe{})
	Collection := Session.DB("heroku_app37083199").C("players")

	var players Players

	if err := Collection.Find(nil).All(&players); err != nil {
		return nil, err
	}

	return &players, nil
}
