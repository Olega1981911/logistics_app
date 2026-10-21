package models

import "time"

// FuelRecord — запись о заправке.
// Одна запись = один чек на АЗС.
type FuelRecord struct {
	ID       int64 `json:"id" db:"id"`
	TenantID int64 `json:"tenant_id" db:"tenant_id"`

	// Кто и на чём
	VehicleID int64  `json:"vehicle_id" db:"vehicle_id"`
	DriverID  *int64 `json:"driver_id,omitempty" db:"driver_id"`
	RouteID   *int64 `json:"route_id,omitempty" db:"route_id"`

	// Что залили
	FuelType      string  `json:"fuel_type" db:"fuel_type"` // diesel, petrol_95, petrol_98, gas, electric
	Liters        float64 `json:"liters" db:"liters"`
	PricePerLiter float64 `json:"price_per_liter" db:"price_per_liter"`
	TotalCost     float64 `json:"total_cost" db:"total_cost"`

	// Показания одометра — ключ для расчёта расхода (л/100км)
	Odometer float64 `json:"odometer" db:"odometer"`

	// Где заправились (опционально)
	StationName string   `json:"station_name" db:"station_name"`
	Lat         *float64 `json:"lat,omitempty" db:"lat"`
	Lng         *float64 `json:"lng,omitempty" db:"lng"`

	// Когда
	RecordedAt time.Time `json:"recorded_at" db:"recorded_at"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
