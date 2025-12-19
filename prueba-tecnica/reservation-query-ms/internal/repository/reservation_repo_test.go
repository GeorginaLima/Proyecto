package repository

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllReservations(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock open: %v", err)
	}
	defer db.Close()
	DB = db

	rows := sqlmock.NewRows([]string{"id", "customer_name", "date", "status"}).
		AddRow("id-1", "Alice", time.Now().UTC(), "CREATED").
		AddRow("id-2", "Bob", time.Now().UTC(), "CREATED")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, customer_name, date, status FROM reservations")).WillReturnRows(rows)

	res, err := GetAllReservations(context.Background())
	if err != nil {
		t.Fatalf("GetAllReservations err: %v", err)
	}
	if len(res) != 2 {
		t.Fatalf("expected 2 rows got %d", len(res))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
