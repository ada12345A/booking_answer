package hotel_availibility

import "testing"

func TestA(t *testing.T) {
	hotels := make(map[string]int)
	reservations := []Reservation{}

	// mock data by the request
	hotels = addHotel("WilliamHotel", 4, hotels)
	hotels = addHotel("RomaHotel", 10, hotels)

	reservations = addReservation(1, 2, "WilliamHotel", reservations)
	reservations = addReservation(1, 4, "RomaHotel", reservations)
	reservations = addReservation(2, 5, "RomaHotel", reservations)

}
