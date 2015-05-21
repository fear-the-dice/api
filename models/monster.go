package models

import (
	"encoding/json"
	"fmt"
	"os"
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
