package models

import "time"

type VehicleStatus string

const (
	VehicleStatusAvailable   VehicleStatus = "available"
	VehicleStatusOnRoute     VehicleStatus = "on_route"
	VehicleStatusMaintenance VehicleStatus = "maintenance"
)

type Vehicle struct {
	Id            uint          `json:"id"`
	LicensePlate  string        `json:"license_plate"` //госномер
	Brand         string        `json:"brand"`
	Model         string        `json:"model"`
	Capacity      float64       `json:"capacity"`                                                        //грузоподъёмность в кг
	Volume        *float64      `json:"volume"`                                                          //объём кузова в м³ (опционально)
	FuelType      string        `json:"fuel_type"`                                                       //тип топлева
	EmissionClass string        `json:"emission_class"`                                                  // например: евро5, евро6
	Status        VehicleStatus `json:"status" validate:"required,oneof=available on_route maintenance"` //статус: "available", "on_route", "maintenance"
	DriverID      *uint         `json:"driver_id"`                                                       //водитель
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
