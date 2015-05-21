package models

// User stores a users name
type Player struct {
	ID         string `json:"ID"`
	Name       string `json:"name"`
	Character  string `json:"character"`
	Initiative int    `json:"initiative"`
	AC         int    `json:"ac"`
	HP         int    `json:"hp"`
	Health     int    `json:"health"`
	Speed      string `json:"speed"`
	Damage     int    `json:"damage"`
}
