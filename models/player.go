package models

import (
	"gopkg.in/mgo.v2"
)

// User stores a users name
type Player struct {
	ID         string `bson:"id" json:"id"`
	Initiative int    `bson:"initiative"`
	AC         int    `bson:"ac"`
	HP         int    `bson:"hp"`
	Health     int    `bson:"health"`
	Damage     int    `bson:"damage"`
	Speed      string `bson:"speed"`
	Name       string `bson:"name"`
	Character  string `bson:"character"`
}
type Players []*Player

func NewPlayer() *Player {
	return &Player{}
}

func FindPlayer(id string) (*Player, error) {
	Setup()

	Session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	defer Session.Close()

	Session.SetSafe(&mgo.Safe{})
	Collection := Session.DB("heroku_app37083199").C("players")

	var player Player

	query := struct {
		id string
	}{
		id,
	}

	if err := Collection.Find(query).One(&player); err != nil {
		return nil, err
	}

	return &player, nil
}

func PopulatePlayers() (*Players, error) {
	Setup()

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
