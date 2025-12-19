package service

import (
	"reservation-query-ms/internal/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestServiceGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock open: %v", err)
	}
	defer db.Close()
	repository.DB = db

	now := time.Now().UTC()
	mock.ExpectQuery("SELECT id, customer_name, date, status FROM reservations").WillReturnRows(sqlmock.NewRows([]string{"id", "customer_name", "date", "status"}).AddRow("id-1", "A", now, "CREATED"))

	svc := NewReservationService()
	_, err = svc.GetAll()
	if err != nil {
		t.Fatalf("GetAll err: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
