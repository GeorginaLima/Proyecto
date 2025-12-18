package model

type Reservation struct {
	ID           string `json:"id"`
	CustomerName string `json:"customer_name"`
	Date         string `json:"date"`
	Status       string `json:"status"`
}
