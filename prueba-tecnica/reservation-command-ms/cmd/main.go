package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"reservation-command-ms/internal/handler"
	"reservation-command-ms/internal/middleware"
	"reservation-command-ms/internal/repository"
	"reservation-command-ms/internal/service"
)

func main() {
	reservationService := service.NewReservationService()
	reservationHandler := handler.NewReservationHandler(reservationService)

	// Initialize DB (CockroachDB/Postgres-compatible). Uses COCKROACH_DSN env var.
	if err := repository.InitDBFromEnv(); err != nil {
		log.Fatal("database init:", err)
	}
	defer repository.CloseDB()

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

	// Cargar certificado del servidor
	cert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
	if err != nil {
		log.Fatal(err)
	}

	// Cargar CA
	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Configuración TLS con mTLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	// Servidor HTTPS
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	log.Println("Command MS running with mTLS on port 8443")
	log.Fatal(server.ListenAndServeTLS("", ""))

}
