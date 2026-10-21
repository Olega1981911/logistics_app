package models

import "time"

// Driver — операционные данные водителя.
// Отдельно от User: User = аутентификация, Driver = права/статус/машина.
type Driver struct {
	ID       int64 `json:"id" db:"id"`
	TenantID int64 `json:"tenant_id" db:"tenant_id"`

	// Связь с User (nullable — водитель может быть без логина)
	UserID *int64 `json:"user_id,omitempty" db:"user_id"`

	// ФИО и контакты
	FullName string `json:"full_name" db:"full_name"`
	Phone    string `json:"phone" db:"phone"`

	// Водительские права — важны для комплаенса (ИИ-агент будет напоминать о продлении)
	LicenseNumber  string    `json:"license_number" db:"license_number"`
	LicenseValidTo time.Time `json:"license_valid_to" db:"license_valid_to"`

	// Статус: available, on_route, resting, offline
	Status string `json:"status" db:"status"`

	// Текущая машина (nullable — если водитель без машины или отдыхает)
	CurrentVehicleID *int64 `json:"current_vehicle_id,omitempty" db:"current_vehicle_id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
