package booking

type MemoryStore struct {
	// seats -> bookings
	bookings map[string]Booking // fast access O(1)

}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	// access the map, seat taken -> error, else add to map
	if _, exists := s.bookings[b.SeatNumber]; exists {
		return ErrSeatAlreadyBooked
	}

	s.bookings[b.SeatNumber] = b
	return nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {
	var result []Booking

	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	
	return result
}