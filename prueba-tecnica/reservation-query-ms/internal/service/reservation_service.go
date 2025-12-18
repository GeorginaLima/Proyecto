package service

import "reservation-query-ms/internal/model"

type ReservationService struct {
	reservations []model.Reservation
}

func NewReservationService() *ReservationService {
	return &ReservationService{
		reservations: []model.Reservation{
			{
				ID:           "example-id",
				CustomerName: "Georgina",
				Date:         "2025-12-20T00:00:00Z",
				Status:       "CREATED",
			},
		},
	}
}

func (s *ReservationService) GetAll() []model.Reservation {
	return s.reservations
}
