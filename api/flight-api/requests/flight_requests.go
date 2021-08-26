package requests

import "time"

type FlightRequest struct {
	Slut string    `json:"slut"`
	Name string    `json:"name"`
	From string    `json:"from"`
	To   string    `json:"to"`
	Date time.Time `json:"date"`
}
