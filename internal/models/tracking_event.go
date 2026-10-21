package models

import "time"

// TrackingEvent — дискретное событие по заказу (или маршрутной точке).
// В отличие от LocationPoint (непрерывный поток GPS), это точечные статусы.
type TrackingEvent struct {
	ID       int64 `json:"id" db:"id"`
	TenantID int64 `json:"tenant_id" db:"tenant_id"`

	// К чему относится: order_id или route_point_id (один из двух)
	OrderID      *int64 `json:"order_id,omitempty" db:"order_id"`
	RoutePointID *int64 `json:"route_point_id,omitempty" db:"route_point_id"`

	// Тип события: created, assigned, picked_up, in_transit, arrived, delivered, exception, returned
	EventType string `json:"event_type" db:"event_type"`

	// Человекочитаемое описание (для клиента и истории)
	Description string `json:"description" db:"description"`

	// Где произошло (опционально — если событие не привязано к GPS-точке)
	Lat *float64 `json:"lat,omitempty" db:"lat"`
	Lng *float64 `json:"lng,omitempty" db:"lng"`

	// Кто инициировал: driver, dispatcher, system, client
	ActorType string `json:"actor_type" db:"actor_type"`
	ActorID   *int64 `json:"actor_id,omitempty" db:"actor_id"`

	// Когда произошло
	OccurredAt time.Time `json:"occurred_at" db:"occurred_at"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
