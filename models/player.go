package models

// User stores a users name
type Monster struct {
	ID         string `json:"ID"`
	Name       string `json:"name"`
	Initiative int    `json:"initiative"`
	AC         int    `json:"ac"`
	HP         int    `json:"hp"`
	Health     int    `json:"health"`
	Speed      string `json:"speed"`
	Damage     int    `json:"damage"`
	Challange  int    `json:"challange"`
	XP         int    `json:"xp"`
	Manual     int    `json:"manual"`
}
