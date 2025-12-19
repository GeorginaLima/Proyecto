package handler

import (
	"encoding/json"
	"net/http"
	"reservation-command-ms/internal/model"
	"reservation-command-ms/internal/service"
)

type ReservationHandler struct {
	service *service.ReservationService
}

func NewReservationHandler(s *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{service: s}
}

func (h *ReservationHandler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var reservation model.Reservation
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	created, err := h.service.CreateReservation(reservation)
	if err != nil {
		http.Error(w, "could not create reservation", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}
func (h *ReservationHandler) UpdateReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	var reservation model.Reservation
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updated, ok, err := h.service.UpdateReservation(id, reservation)
	if err != nil {
		http.Error(w, "could not update reservation", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (h *ReservationHandler) DeleteReservation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	ok, err := h.service.DeleteReservation(id)
	if err != nil {
		http.Error(w, "could not delete reservation", http.StatusInternalServerError)
		return
	}
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
