package model

type Cinema struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Rate     float32 `json:"rate"`
}
