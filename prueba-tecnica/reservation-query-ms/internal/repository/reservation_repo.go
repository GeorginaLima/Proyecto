package repository

import (
	"context"
	"fmt"
	"time"

	"reservation-query-ms/internal/model"
)

func GetAllReservations(ctx context.Context) ([]model.Reservation, error) {
	if DB == nil {
		return nil, fmt.Errorf("db not initialized")
	}

	rows, err := DB.QueryContext(ctx, `SELECT id, customer_name, date, status FROM reservations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []model.Reservation
	for rows.Next() {
		var id, customerName, status string
		var t time.Time
		if err := rows.Scan(&id, &customerName, &t, &status); err != nil {
			return nil, err
		}
		res = append(res, model.Reservation{
			ID:           id,
			CustomerName: customerName,
			Date:         t.Format(time.RFC3339),
			Status:       status,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
