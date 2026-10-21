package models

import "time"

// Subscription — активная подписка компании на тариф.
// Один Tenant может иметь одну активную подписку в момент времени.
type Subscription struct {
	ID       int64 `json:"id" db:"id"`
	TenantID int64 `json:"tenant_id" db:"tenant_id"`
	PlanID   int64 `json:"plan_id" db:"plan_id"`

	// Статус: trialing, active, past_due, canceled, suspended
	Status string `json:"status" db:"status"`

	// Начало и конец текущего оплаченного периода
	CurrentPeriodStart time.Time `json:"current_period_start" db:"current_period_start"`
	CurrentPeriodEnd   time.Time `json:"current_period_end" db:"current_period_end"`

	// Когда подписка была отменена (nil = не отменена)
	CanceledAt *time.Time `json:"canceled_at,omitempty" db:"canceled_at"`

	// Внешний ID в платёжной системе (Stripe, Paddle и т.д.)
	ExternalSubscriptionID *string `json:"external_subscription_id,omitempty" db:"external_subscription_id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
