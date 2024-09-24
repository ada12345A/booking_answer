package hotel_availibility

import "fmt"

/**
 * You are building a small command-line application to calculate hotel availability for a city. Your application reads in two (2) data files, and outputs its answer to STDOUT.
*Your application will read in:
*· a list of hotels along with how many rooms each contains (in no particular order)
*· a list of bookings that have been made (in no particular order)
*Your application will then print the list of all hotels which have availability for check-in and check- out date range, if any.
*Do not worry about whether a specific room is available in a hotel for the entire booking period without switching rooms: availability is defined as the hotel having at least one (1) available room for each night of the target stay, regardless of whether it's the same room from day to day.
*
*Data Files
*hotels.csv
*# Name, Rooms
*Westin, 10
*Best Western, 20
*Hilton, 10
*...
*
*bookings.csv
*# Name, Checkin, Checkout
*Hilton, 2015-04-02, 2015-04-03
*Hilton, 2015-04-02, 2015-04-04
*Westin, 2015-05-01, 2015-05-20
 *
*/

// https://github.com/zhaohuabing/booking-interview-practice/blob/master/src/com/zhaohuabing/HotelAvailability.java
type Reservation struct {
	startDate int
	endDate   int
	hotel     string
}

func isConflict(startDate2, endDate2 int, r Reservation) bool {
	// i think this implementation is wrong, so i changed
	/*if r.startDate >= startDate2 && r.endDate <= endDate2 {
		return true
	}
	return false
	*/

	// should be like below
	// leave at AM and check in at PM
	if r.endDate <= startDate2 || r.startDate >= endDate2 {
		return false
	}
	return true
}

func addReservation(startDate int, endDate int, hotel string, reservations []Reservation) []Reservation {
	reservations = append(reservations, Reservation{
		startDate: startDate,
		endDate:   endDate,
		hotel:     hotel,
	})
	return reservations
}

func addHotel(name string, room int, hotel map[string]int) map[string]int {
	hotel[name] = room
	return hotel
}

func solve(startDate, endDate int, reservations []Reservation, hotel map[string]int) {
	for i := 0; i < len(reservations); i++ {
		r := reservations[i]
		if isConflict(startDate, endDate, r) {
			hotel[r.hotel] -= 1
		}
	}

	for hotelName, leftRoomCnt := range hotel {
		fmt.Print("Available rooms of %v:%v", hotelName, leftRoomCnt)
	}
}
