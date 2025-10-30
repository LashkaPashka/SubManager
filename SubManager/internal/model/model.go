package model

import "time"

type SubscriptionInputModel struct {
	ServiceName string    `json:"service_name"`
	Price       uint      `json:"price"`
	UserID      string    `json:"user_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}