package models

import (
	"gopkg.in/mgo.v2"
)

// User stores a users name
type (
	Monster struct {
		ID         string `json:"id"`
		Initiative int    `json:"initiative"`
		AC         int    `json:"ac"`
		HP         int    `json:"hp"`
		Health     int    `json:"health"`
		Damage     int    `json:"damage"`
		Speed      string `json:"speed"`
		Monster    string `json:"monster"`
		Challange  int    `json:"challange"`
		XP         int    `json:"xp"`
		Manual     int    `json:"manual"`
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

func FindMonster(id string) (*Monster, error) {
	Setup()

	Session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	defer Session.Close()

	Session.SetSafe(&mgo.Safe{})
	Collection := Session.DB("heroku_app37083199").C("monsters")

	var monster Monster

	query := struct {
		id string
	}{
		id,
	}

	if err := Collection.Find(query).One(&monster); err != nil {
		return nil, err
	}

	return &monster, nil
}

func PopulateMonsters() (*Monsters, error) {
	Setup()

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
