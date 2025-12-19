package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"reservation-command-ms/internal/model"
)

func CreateReservation(ctx context.Context, r model.Reservation) (model.Reservation, error) {
	if DB == nil {
		return model.Reservation{}, fmt.Errorf("db not initialized")
	}

	if r.ID == "" {
		// keep caller to set ID if desired
	}

	_, err := DB.ExecContext(ctx,
		`INSERT INTO reservations (id, customer_name, date, status) VALUES ($1, $2, $3, $4)`,
		r.ID, r.CustomerName, r.Date, r.Status,
	)
	if err != nil {
		return model.Reservation{}, err
	}
	return r, nil
}

func UpdateReservation(ctx context.Context, id string, r model.Reservation) (model.Reservation, bool, error) {
	if DB == nil {
		return model.Reservation{}, false, fmt.Errorf("db not initialized")
	}
	res, err := DB.ExecContext(ctx,
		`UPDATE reservations SET customer_name = $1, date = $2, status = $3 WHERE id = $4`,
		r.CustomerName, r.Date, r.Status, id,
	)
	if err != nil {
		return model.Reservation{}, false, err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return model.Reservation{}, false, nil
	}
	r.ID = id
	return r, true, nil
}

func DeleteReservation(ctx context.Context, id string) (bool, error) {
	if DB == nil {
		return false, fmt.Errorf("db not initialized")
	}
	res, err := DB.ExecContext(ctx, `DELETE FROM reservations WHERE id = $1`, id)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}

func GetReservationByID(ctx context.Context, id string) (model.Reservation, bool, error) {
	if DB == nil {
		return model.Reservation{}, false, fmt.Errorf("db not initialized")
	}
	var r model.Reservation
	var t time.Time
	row := DB.QueryRowContext(ctx, `SELECT id, customer_name, date, status FROM reservations WHERE id = $1`, id)
	if err := row.Scan(&r.ID, &r.CustomerName, &t, &r.Status); err != nil {
		if err == sql.ErrNoRows {
			return model.Reservation{}, false, nil
		}
		return model.Reservation{}, false, err
	}
	r.Date = t
	return r, true, nil
}
