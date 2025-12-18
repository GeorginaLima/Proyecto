package model

import "time"

type Reservation struct {
	ID           string    `json:"id"`
	CustomerName string    `json:"customer_name"`
	Date         time.Time `json:"date"`
	Status       string    `json:"status"`
}
