package booking

import "sync"

type ConcurrentStore struct {
	// seats -> bookings
	bookings map[string]Booking // fast access O(1)
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (s *ConcurrentStore) Book(b Booking) error {

	s.Lock()

	defer s.Unlock()

	// access the map, seat taken -> error, else add to map
	if _, exists := s.bookings[b.SeatNumber]; exists {
		return ErrSeatAlreadyBooked
	}

	s.bookings[b.SeatNumber] = b
	return nil
}

func (s *ConcurrentStore) ListBookings(movieID string) []Booking {
	
	s.RLock()
	defer s.RUnlock()

	var result []Booking

	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	
	return result
}