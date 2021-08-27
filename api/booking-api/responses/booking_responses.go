package responses

type BookingResponse struct {
	Id          string `json:"id"`
	CustomerId  string `json:"customer_id"`
	FlightId    string `json:"flight_id"`
	Code        string `json:"code"`
	Status      string `json:"status"`
	Booked_date string `json:"booked_date"`
}

type BookingCancelResponse struct {
	Id         string `json:"id"`
	CustomerId string `json:"customer_id"`
	FlightId   string `json:"flight_id"`
	Code       string `json:"code"`
	Status     string `json:"status"`
	CancelDate string `json:"cancel_date"`
}

type BookingViewResponse struct {
	Code        string           `json:"booking_code"`
	Booked_date string           `json:"booking_date"`
	Customer    CustomerResponse `json:"customer"`
	Flight      FlightResponse   `json:"flight"`
}

type CustomerResponse struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type FlightResponse struct {
	Id         string `json:"id"`
	Name       string `json:"flight_plane"`
	From       string `json:"from"`
	To         string `json:"to"`
	Status     string `json:"status"`
	Slot       string `json:"slot"`
	DepartDate string `json:"depart_date"`
	DepartTime string `json:"depart_time"`
}
