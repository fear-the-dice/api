package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// User stores a users name
type (
	Player struct {
		ID         string `json:"ID"`
		Name       string `json:"name"`
		Character  string `json:"character"`
		Initiative int    `json:"initiative"`
		AC         int    `json:"ac"`
		HP         int    `json:"hp"`
		Health     int    `json:"health"`
		Damage     int    `json:"damage"`
		Speed      int    `json:"speed"`
	}
	Players []*Player
)

func NewPlayer() *Player {
	return &Player{}
}

func PopulatePlayers() (*Players, error) {
	var players Players
	playersJSON, err := os.Open("players.json")
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(playersJSON)
	if err = jsonParser.Decode(&players); err != nil {
		fmt.Printf("%s", err.Error())
	}

	return &players, nil
}
