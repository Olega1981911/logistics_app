package models

import "time"

type Payment struct {
	Id          uint      `json:"id"`
	InvoiceID   uint      `json:"invoice_id"`
	Amount      int       `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Method      string    `json:"method"`    //"cash", "bank_transfer", "credit_card" (enum)
	Reference   *string   `json:"reference"` //номер транзакции назначения платежа
	Status      string    `json:"status"`    //"pending", "completed", "failed"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
