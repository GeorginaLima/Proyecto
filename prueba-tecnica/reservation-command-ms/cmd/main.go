package main

import (
	"log"
	"net/http"

	"reservation-command-ms/internal/handler"
	"reservation-command-ms/internal/middleware"
	"reservation-command-ms/internal/service"
)

func main() {
	reservationService := service.NewReservationService()
	reservationHandler := handler.NewReservationHandler(reservationService)

	// RUTAS
	//http.HandleFunc("/create-reservation", reservationHandler.CreateReservation)
	// CREATE
	http.HandleFunc(
		"/create-reservation",
		middleware.JWTMiddleware(reservationHandler.CreateReservation),
	)

	// UPDATE
	http.HandleFunc(
		"/update-reservation",
		middleware.JWTMiddleware(reservationHandler.UpdateReservation),
	)

	// DELETE
	http.HandleFunc(
		"/delete-reservation",
		middleware.JWTMiddleware(reservationHandler.DeleteReservation),
	)

	log.Println("Command MS running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
