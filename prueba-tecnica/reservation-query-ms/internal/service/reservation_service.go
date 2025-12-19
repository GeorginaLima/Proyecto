package service

import (
	"context"
	"reservation-query-ms/internal/model"
	"reservation-query-ms/internal/repository"
)

type ReservationService struct{}

func NewReservationService() *ReservationService {
	return &ReservationService{}
}

func (s *ReservationService) GetAll() ([]model.Reservation, error) {
	ctx := context.Background()
	return repository.GetAllReservations(ctx)
}
