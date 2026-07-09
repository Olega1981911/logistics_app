package models

type Cargo struct {
	ID          uint    `json:"primaryKey" json:"primaryKey"`
	ClientId    uint    `json:"client_id"`
	Description string  `json:"description"`
	Weight      float64 `json:"weight"`
	Volume      float64 `json:"volume"`
}
