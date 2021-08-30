package requests

type BookingRequest struct {
	Id         string `json:"id"`
	Slut       string `json:"slut"`
	CustomerId string `json:"customer_id"`
	FlightId   string `json:"flight_id"`
}

type BookingCodeRequest struct {
	Code string `json:"code"`
}
