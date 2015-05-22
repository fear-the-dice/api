package models

import (
	"os"

	"gopkg.in/mgo.v2"
)

// User stores a users name
type (
	Monster struct {
		Base
		Monster   string `json:"monster"`
		Challange int    `json:"challange"`
		XP        int    `json:"xp"`
		Manual    int    `json:"manual"`
	}
	Monsters []*Monster
)

var (
	Collection *mgo.Collection
	Session    mgo.Session
)

func NewMonster() *Monster {
	return &Monster{}
}

func PopulateMonsters() (*Monsters, error) {
	uri := os.Getenv("MONGOLAB_URI")
	Session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	defer Session.Close()

	Session.SetSafe(&mgo.Safe{})
	Collection := Session.DB("heroku_app37083199").C("monsters")

	var monsters Monsters

	if err := Collection.Find(nil).All(&monsters); err != nil {
		return nil, err
	}

	return &monsters, nil
}
