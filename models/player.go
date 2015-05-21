package models

import (
	"encoding/json"
	"fmt"
	"os"
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
