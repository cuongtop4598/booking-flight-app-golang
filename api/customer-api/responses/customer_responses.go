package responses

import (
	"github.com/cuongtop4598/booking_assigment/booking-flight-app-golang/pb"
)

type BookingHistoryReponse struct {
	BookingCode   string             `json:"booking_code"`
	BookingStatus string             `json:"booking_status"`
	BookedDate    string             `json:"booked_date"`
	Flight        *pb.FlightResponse `json:"flight"`
}

type ListBookingHistoryReponse struct {
	BookingHistorys BookingHistoryReponse `json:"history"`
}
