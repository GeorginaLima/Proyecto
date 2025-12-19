package handler

import (
	"encoding/json"
	"net/http"

	"reservation-query-ms/internal/service"
)

type ReservationHandler struct {
	service *service.ReservationService
}

func NewReservationHandler(s *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{service: s}
}

func (h *ReservationHandler) GetReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	res, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "could not fetch reservations", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
