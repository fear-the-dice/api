package models

// User stores a users name
type (
	Base struct {
		ID         string `json:"ID"`
		Initiative int    `json:"initiative"`
		AC         int    `json:"ac"`
		HP         int    `json:"hp"`
		Health     int    `json:"health"`
		Damage     int    `json:"damage"`
		Speed      int    `json:"speed"`
	}
)
