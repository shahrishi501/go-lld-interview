package booking

import "errors"

var (
	ErrSeatAlreadyBooked = errors.New("seat already booked")
)

// Booking represents a confirmed seat reservation
type Booking struct {
	ID string
	MovieID string
	SeatNumber string
	UserID string
	Status string
}

// For dependency injection, we can define an interface for the booking store
type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}