package responses

import "time"

type FlightReponse struct {
	Id            string    `json:"id"`
	Slut          string    `json:"slut"`
	Name          string    `json:"name"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Date          time.Time `json:"date"`
	Status        string    `json:"status"`
	AvailableSlot int32     `json:"available_slot"`
}
