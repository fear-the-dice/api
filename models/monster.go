package models

import (
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User stores a users name
type (
	Monster struct {
		ID         bson.ObjectId `bson:"_id" json:"id"`
		Initiative int           `json:"initiative"`
		AC         int           `json:"ac"`
		HP         int           `json:"hp"`
		Health     int           `json:"health"`
		Damage     int           `json:"damage"`
		Speed      string        `json:"speed"`
		Monster    string        `json:"monster"`
		Challange  int           `json:"challange"`
		XP         int           `json:"xp"`
		Manual     int           `json:"manual"`
		STR        int           `json:"str"`
		DEX        int           `json:"dex"`
		CON        int           `json:"con"`
		INT        int           `json:"int"`
		WIS        int           `json:"wis"`
		CHA        int           `json:"cha"`
	}
	Monsters []*Monster
)

var (
	Collection *mgo.Collection
	Session    mgo.Session
)

func NewMonster() *Monster {
	id := bson.NewObjectId()
	return &Monster{ID: id}
}

func InsertMonster(monster Monster) (*Monster, error) {
	uri := os.Getenv("MONGOLAB_URI")
	db := os.Getenv("DB")

	Session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	defer Session.Close()

	Collection := Session.DB(db).C("monsters")

	if err := Collection.Insert(monster); err != nil {
		return nil, err
	}

	return &monster, nil
}

func FindMonster(id bson.ObjectId) (*Monster, error) {
	uri := os.Getenv("MONGOLAB_URI")
	db := os.Getenv("DB")

	Session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	defer Session.Close()

	Session.SetSafe(&mgo.Safe{})
	Collection := Session.DB(db).C("monsters")

	var monster Monster

	if err := Collection.Find(bson.M{"_id": id}).One(&monster); err != nil {
		return nil, err
	}

	return &monster, nil
}

func PopulateMonsters() (*Monsters, error) {
	uri := os.Getenv("MONGOLAB_URI")
	db := os.Getenv("DB")

	Session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	defer Session.Close()

	Session.SetSafe(&mgo.Safe{})
	Collection := Session.DB(db).C("monsters")

	var monsters Monsters

	if err := Collection.Find(nil).All(&monsters); err != nil {
		return nil, err
	}

	return &monsters, nil
}

func DeleteMonster(id bson.ObjectId) error {
	uri := os.Getenv("MONGOLAB_URI")
	db := os.Getenv("DB")

	Session, err := mgo.Dial(uri)
	if err != nil {
		return nil
	}
	defer Session.Close()

	Collection := Session.DB(db).C("monsters")

	if err := Collection.Remove(bson.M{"_id": id}); err != nil {
		return err
	}

	return nil
}

func UpdateMonster(id bson.ObjectId, monster Monster) error {
	uri := os.Getenv("MONGOLAB_URI")
	db := os.Getenv("DB")

	Session, err := mgo.Dial(uri)
	if err != nil {
		return nil
	}
	defer Session.Close()

	Collection := Session.DB(db).C("monsters")

	_, err = Collection.Upsert(bson.M{"_id": id}, monster)
	if err != nil {
		return err
	}

	return nil
}
