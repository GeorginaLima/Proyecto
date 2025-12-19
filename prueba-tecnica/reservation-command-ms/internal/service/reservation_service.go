package service

import (
	"context"
	"reservation-command-ms/internal/model"
	"reservation-command-ms/internal/repository"

	"github.com/google/uuid"
)

type ReservationService struct{}

func NewReservationService() *ReservationService {
	return &ReservationService{}
}

func (s *ReservationService) CreateReservation(r model.Reservation) (model.Reservation, error) {
	r.ID = uuid.New().String()
	r.Status = "CREATED"
	ctx := context.Background()
	return repository.CreateReservation(ctx, r)
}

func (s *ReservationService) UpdateReservation(id string, r model.Reservation) (model.Reservation, bool, error) {
	ctx := context.Background()
	return repository.UpdateReservation(ctx, id, r)
}

func (s *ReservationService) DeleteReservation(id string) (bool, error) {
	ctx := context.Background()
	return repository.DeleteReservation(ctx, id)
}
