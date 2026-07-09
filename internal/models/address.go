package models

type Address struct {
	Id          uint   `json:"id"`
	OwnerType   string `json:"owner_type"` // "client", "delivery_point"
	OwnerID     uint   `json:"owner_id"`
	Type        string `json:"type"` // "legal", "actual", "delivery" (для клиента); "pickup", "destination" (для доставки)
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	PostalCode  string `json:"postal_code"`
	City        string `json:"city"`
	Country     string `json:"country"`
	IsDefault   bool   `json:"is_default"` // флаг основного адреса (если нужно)
}
