package repository

import (
	"context"
	"saloonbook/internal/models"
	"sync"
	"sync/atomic"
)

type BookingRepository struct {
	mu       sync.RWMutex
	bookings map[int]models.Booking
	seq      int32
}

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		bookings: make(map[int]models.Booking),
	}
}

func (r *BookingRepository) Create(ctx context.Context, booking models.Booking) (*models.Booking, error) {
	id := int(atomic.AddInt32(&r.seq, 1))
	booking.ID = id
	booking.Status = "confirmed"

	r.mu.Lock()
	r.bookings[id] = booking
	r.mu.Unlock()

	return &booking, nil
}

func (r *BookingRepository) List(ctx context.Context) ([]models.Booking, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]models.Booking, 0, len(r.bookings))
	for _, b := range r.bookings {
		result = append(result, b)
	}
	return result, nil
}

func (r *BookingRepository) GetByID(ctx context.Context, id int) (*models.Booking, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if booking, ok := r.bookings[id]; ok {
		return &booking, nil
	}
	return nil, nil
}
