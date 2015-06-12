package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User stores a users name
type (
	Player struct {
		ID         bson.ObjectId `bson:"_id" json:"id"`
		Initiative int           `json:"initiative"`
		AC         int           `json:"ac"`
		HP         int           `json:"hp"`
		Health     int           `json:"health"`
		Damage     int           `json:"damage"`
		Speed      string        `json:"speed"`
		Name       string        `json:"name"`
		Character  string        `json:"character"`
		Thumb      string        `json:"thumb"`
	}
	Players []*Player
)

func NewPlayer() *Player {
	id := bson.NewObjectId()
	return &Player{ID: id}
}

func InsertPlayer(player Player) (*Player, error) {
	Session, err := mgo.Dial(dbOptions.Host)
	if err != nil {
		return nil, err
	}
	defer Session.Close()

	Collection := Session.DB(dbOptions.Database).C("players")

	if err := Collection.Insert(player); err != nil {
		return nil, err
	}

	return &player, nil
}

func FindPlayer(id bson.ObjectId) (*Player, error) {
	Session, err := mgo.Dial(dbOptions.Host)
	if err != nil {
		return nil, err
	}
	defer Session.Close()

	Session.SetSafe(&mgo.Safe{})
	Collection := Session.DB(dbOptions.Database).C("players")

	var player Player

	if err := Collection.FindId(id).One(&player); err != nil {
		return nil, err
	}

	return &player, nil
}

func PopulatePlayers() (*Players, error) {
	Session, err := mgo.Dial(dbOptions.Host)
	if err != nil {
		return nil, err
	}
	defer Session.Close()

	Session.SetSafe(&mgo.Safe{})
	Collection := Session.DB(dbOptions.Database).C("players")

	var players Players

	if err := Collection.Find(nil).All(&players); err != nil {
		return nil, err
	}

	return &players, nil
}

func DeletePlayer(id bson.ObjectId) error {
	Session, err := mgo.Dial(dbOptions.Host)
	if err != nil {
		return nil
	}
	defer Session.Close()

	Collection := Session.DB(dbOptions.Database).C("players")

	if err := Collection.Remove(bson.M{"_id": id}); err != nil {
		return err
	}

	return nil
}

func UpdatePlayer(id bson.ObjectId, player Player) error {
	Session, err := mgo.Dial(dbOptions.Host)
	if err != nil {
		return nil
	}
	defer Session.Close()

	Collection := Session.DB(dbOptions.Database).C("players")

	_, err = Collection.Upsert(bson.M{"_id": id}, player)
	if err != nil {
		return err
	}

	return nil
}
