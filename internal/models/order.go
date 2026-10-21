package models

import "time"

// Order — заказ на доставку груза от точки A до точки B.
// Жизненный цикл: created → assigned → in_transit → delivered (или canceled).
type Order struct {
	ID       int64 `json:"id" db:"id"`
	TenantID int64 `json:"tenant_id" db:"tenant_id"` // чей это заказ (SaaS)
	ClientID int64 `json:"client_id" db:"client_id"` // кто заказал

	// Откуда и куда — ссылки на таблицу addresses
	PickupAddressID   int64 `json:"pickup_address_id" db:"pickup_address_id"`
	DeliveryAddressID int64 `json:"delivery_address_id" db:"delivery_address_id"`

	// Что везём — ссылка на таблицу cargo
	CargoID int64 `json:"cargo_id" db:"cargo_id"`

	// Состояние заказа: created, assigned, in_transit, delivered, canceled
	Status string `json:"status" db:"status"`

	// Приоритет: 1 = низкий, 5 = срочный. Влияет на сортировку в диспетчерской.
	Priority int `json:"priority" db:"priority"`

	// Когда клиент хочет pickup и delivery (может быть null, если «как можно скорее»)
	RequestedPickupAt   *time.Time `json:"requested_pickup_at,omitempty" db:"requested_pickup_at"`
	RequestedDeliveryAt *time.Time `json:"requested_delivery_at,omitempty" db:"requested_delivery_at"`

	// Фактическое время — заполняется по мере выполнения
	ActualPickupAt   *time.Time `json:"actual_pickup_at,omitempty" db:"actual_pickup_at"`
	ActualDeliveryAt *time.Time `json:"actual_delivery_at,omitempty" db:"actual_delivery_at"`

	// Кому назначен — NULL, пока заказ не распределён
	AssignedVehicleID *int64 `json:"assigned_vehicle_id,omitempty" db:"assigned_vehicle_id"`
	AssignedDriverID  *int64 `json:"assigned_driver_id,omitempty" db:"assigned_driver_id"`

	// Цена заказа (что заплатит клиент). Себестоимость считается отдельно в service.
	Price float64 `json:"price" db:"price"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
