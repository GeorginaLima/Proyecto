package repository

import (
	"context"
	"regexp"
	"testing"
	"time"

	"reservation-command-ms/internal/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateReservation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()
	DB = db

	r := model.Reservation{
		ID:           "id-1",
		CustomerName: "Alice",
		Date:         time.Now(),
		Status:       "CREATED",
	}

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO reservations (id, customer_name, date, status) VALUES ($1, $2, $3, $4)")).
		WithArgs(r.ID, r.CustomerName, sqlmock.AnyArg(), r.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	got, err := CreateReservation(ctx, r)
	if err != nil {
		t.Fatalf("CreateReservation error: %v", err)
	}
	if got.ID != r.ID {
		t.Fatalf("expected id %s got %s", r.ID, got.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestUpdateReservation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()
	DB = db

	id := "id-1"
	r := model.Reservation{
		CustomerName: "Bob",
		Date:         time.Now(),
		Status:       "UPDATED",
	}

	mock.ExpectExec(regexp.QuoteMeta("UPDATE reservations SET customer_name = $1, date = $2, status = $3 WHERE id = $4")).
		WithArgs(r.CustomerName, sqlmock.AnyArg(), r.Status, id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := context.Background()
	got, ok, err := UpdateReservation(ctx, id, r)
	if err != nil {
		t.Fatalf("UpdateReservation error: %v", err)
	}
	if !ok {
		t.Fatalf("expected ok true")
	}
	if got.ID != id {
		t.Fatalf("expected id %s got %s", id, got.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestDeleteReservation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()
	DB = db

	id := "id-1"
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM reservations WHERE id = $1")).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	ok, err := DeleteReservation(context.Background(), id)
	if err != nil {
		t.Fatalf("DeleteReservation error: %v", err)
	}
	if !ok {
		t.Fatalf("expected deleted true")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestGetReservationByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()
	DB = db

	id := "id-1"
	customer := "Carol"
	dt := time.Now().UTC()

	rows := sqlmock.NewRows([]string{"id", "customer_name", "date", "status"}).
		AddRow(id, customer, dt, "CREATED")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, customer_name, date, status FROM reservations WHERE id = $1")).
		WithArgs(id).
		WillReturnRows(rows)

	r, ok, err := GetReservationByID(context.Background(), id)
	if err != nil {
		t.Fatalf("GetReservationByID error: %v", err)
	}
	if !ok {
		t.Fatalf("expected found true")
	}
	if r.CustomerName != customer {
		t.Fatalf("expected customer %s got %s", customer, r.CustomerName)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
