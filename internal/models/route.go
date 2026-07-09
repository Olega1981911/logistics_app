package models

import "time"

type RouteStatus string

const (
	RouteStatusPlanned    RouteStatus = "planned"
	RouteStatusInProgress RouteStatus = "in_progress"
	RouteStatusCompleted  RouteStatus = "completed"
	RouteStatusCancelled  RouteStatus = "cancelled"
)

type PointType string

const (
	PointTypePickup   PointType = "pickup"
	PointTypeDelivery PointType = "delivery"
	PointTypeStop     PointType = "stop"
)

type PointStatus string

const (
	PointStatusPending  PointStatus = "pending"
	PointStatusArrived  PointStatus = "arrived"
	PointStatusDeparted PointStatus = "departed"
)

type Route struct {
	Id            uint        `json:"id"`
	CargoId       uint        `json:"cargo_id"`
	VehicleId     uint        `json:"vehicle_id"`
	DriverId      uint        `json:"driver_id"`
	Status        RouteStatus `json:"status" validate:"required,oneof=planned in_progress completed cancelled"` //"planned", "in_progress", "completed", "cancelled"
	TotalDistance float64     `json:"total_distance"`                                                           // km
	StartedAt     *time.Time  `json:"started_at"`
	CompletedAt   *time.Time  `json:"completed_at"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

type RoutePoint struct {
	Id             uint        `json:"id"`
	RouteId        uint        `json:"route_id"`
	Type           PointType   `json:"type"` //"pickup", "delivery", "stop" (enum)
	AddressID      uint        `json:"address_id"`
	Sequence       int         `json:"sequence"`
	PlannedArrival time.Time   `json:"planned_arrival"` //плановое время прибытия
	ActualArrival  *time.Time  `json:"actual_arrival"`  //фактическое время прибытия
	Status         PointStatus `json:"status"`          //"pending", "arrived", "departed"
}
