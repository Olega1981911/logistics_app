package models

import "time"

// LocationPoint — одна GPS-точка от машины.
// Append-only: только INSERT, без UPDATE и DELETE.
// В БД партиционируется по RecordedAt (месяц/день).
type LocationPoint struct {
	ID        int64 `json:"id" db:"id"`
	TenantID  int64 `json:"tenant_id" db:"tenant_id"`
	VehicleID int64 `json:"vehicle_id" db:"vehicle_id"`

	// Координаты — float64 даёт ~15 значащих цифр, точность ~1 мм
	Lat float64 `json:"lat" db:"lat"`
	Lng float64 `json:"lng" db:"lng"`

	// Высота над уровнем моря — нужна для маршрутов с учётом рельефа
	Elevation float64 `json:"elevation" db:"elevation"`

	// Скорость (км/ч) и направление (градусы 0–360)
	Speed   float64 `json:"speed" db:"speed"`
	Heading float64 `json:"heading" db:"heading"`

	// Точность GPS в метрах (чем меньше — тем лучше). Позволяет отбрасывать «плохие» точки.
	Accuracy float64 `json:"accuracy" db:"accuracy"`

	// Когда была снята точка (UTC). Ключевое поле для партиционирования и индексов.
	RecordedAt time.Time `json:"recorded_at" db:"recorded_at"`
}
