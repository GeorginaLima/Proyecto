package service

import (
	"testing"
	"time"

	"reservation-command-ms/internal/model"
	"reservation-command-ms/internal/repository"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestServiceCreateUpdateDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock open: %v", err)
	}
	defer db.Close()
	repository.DB = db

	svc := NewReservationService()

	// Create
	mock.ExpectExec("INSERT INTO reservations").WillReturnResult(sqlmock.NewResult(1, 1))
	r := model.Reservation{CustomerName: "A", Date: time.Now(), Status: "CREATED"}
	created, err := svc.CreateReservation(r)
	if err != nil {
		t.Fatalf("create err: %v", err)
	}
	if created.ID == "" {
		t.Fatalf("expected id set")
	}

	// Update
	mock.ExpectExec("UPDATE reservations").WillReturnResult(sqlmock.NewResult(0, 1))
	_, ok, err := svc.UpdateReservation(created.ID, model.Reservation{CustomerName: "B", Date: time.Now(), Status: "UPDATED"})
	if err != nil {
		t.Fatalf("update err: %v", err)
	}
	if !ok {
		t.Fatalf("expected update ok")
	}

	// Delete
	mock.ExpectExec("DELETE FROM reservations").WillReturnResult(sqlmock.NewResult(0, 1))
	delOk, err := svc.DeleteReservation(created.ID)
	if err != nil {
		t.Fatalf("delete err: %v", err)
	}
	if !delOk {
		t.Fatalf("expected delete true")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
