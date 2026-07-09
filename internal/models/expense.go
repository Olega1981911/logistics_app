package models

import "time"

type Expense struct {
	ID             uint      `json:"id"`
	Type           string    `json:"type"`     //"fuel", "maintenance", "toll", "parking", "salary" (enum)
	Amount         int       `json:"amount"`   //сумма в минимальных единицах
	Currency       string    `json:"currency"` //"usd", eur
	ExpenseDate    time.Time `json:"expense_date"`
	Description    *string   `json:"description"`
	VehicleID      *uint     `json:"vehicle_id"`      //ссылка на транспорт (если расход привязан к конкретному авто)
	RouteID        *uint     `json:"route_id"`        //ссылка на маршрут (если расход в рейсе)
	Supplier       *string   `json:"supplier"`        //поставщик услуг (например, заправка)
	DocumentNumber *string   `json:"document_number"` //номер счета/чека
	VatRate        *int      `json:"vat_rate"`        //ставка НДС (если применимо)
	VatAmount      *int      `json:"vat_amount"`      //сумма НДС
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
