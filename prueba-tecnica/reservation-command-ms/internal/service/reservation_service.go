package service

import (
	"reservation-command-ms/internal/model"

	"github.com/google/uuid"
)

type ReservationService struct {
	reservations map[string]model.Reservation
}

func NewReservationService() *ReservationService {
	return &ReservationService{}
}

func (s *ReservationService) CreateReservation(r model.Reservation) model.Reservation {
	r.ID = uuid.New().String()
	r.Status = "CREATED"
	return r
}
func (s *ReservationService) UpdateReservation(id string, r model.Reservation) (model.Reservation, bool) {
	if _, exists := s.reservations[id]; !exists {
		return model.Reservation{}, false
	}

	r.ID = id
	s.reservations[id] = r
	return r, true
}

func (s *ReservationService) DeleteReservation(id string) bool {
	if _, exists := s.reservations[id]; !exists {
		return false
	}

	delete(s.reservations, id)
	return true
}
