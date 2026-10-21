package models

import "time"

type Tenant struct {
	ID           int64     `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	PlanID       int64     `json:"plan_id" db:"plan_id"`
	BillingEmail string    `json:"billing_email" db:"billing_email"`
	Status       string    `json:"status" db:"status"` // active, canceled, suspended
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
