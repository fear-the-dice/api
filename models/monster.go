package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// User stores a users name
type (
	Monster struct {
		ID         string `json:"ID"`
		Monster    string `json:"monster"`
		Initiative int    `json:"initiative"`
		AC         int    `json:"ac"`
		HP         int    `json:"hp"`
		Health     int    `json:"health"`
		Damage     int    `json:"damage"`
		Challange  int    `json:"challange"`
		XP         int    `json:"xp"`
		Manual     int    `json:"manual"`
		Speed      int    `json:"speed"`
	}
	Monsters []*Monster
)

func NewMonster() *Monster {
	return &Monster{}
}

func PopulateMonsters() (*Monsters, error) {
	var monsters Monsters
	monstersJSON, err := os.Open("monsters.json")
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(monstersJSON)
	if err = jsonParser.Decode(&monsters); err != nil {
		fmt.Printf("%s", err.Error())
	}

	return &monsters, nil
}
