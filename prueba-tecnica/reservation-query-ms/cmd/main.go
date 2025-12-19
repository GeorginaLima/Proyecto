package main

import (
	"log"
	"net/http"

	"reservation-query-ms/internal/handler"
	"reservation-query-ms/internal/repository"
	"reservation-query-ms/internal/service"
)

func main() {
	reservationService := service.NewReservationService()
	reservationHandler := handler.NewReservationHandler(reservationService)

	if err := repository.InitDBFromEnv(); err != nil {
		log.Fatal("database init:", err)
	}
	defer repository.CloseDB()

	http.HandleFunc("/reservation", reservationHandler.GetReservation)

	log.Println("Query MS running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
