package models

import "time"

type InvoiceStatus string

const (
	InvoiceStatusDraft   InvoiceStatus = "draft"
	InvoiceStatusSent    InvoiceStatus = "sent"
	InvoiceStatusPaid    InvoiceStatus = "paid"
	InvoiceStatusOverdue InvoiceStatus = "overdue"
)

type Invoice struct {
	ID          uint          `json:"id"`
	Number      string        `json:"number"`                                                   // type: "INV-2025-001"
	ClientID    uint          `json:"client_id"`                                                //плательщик
	IssueDate   time.Time     `json:"issue_date"`                                               //дата выставления
	DueDate     time.Time     `json:"due_date"`                                                 // срок оплаты
	TotalAmount int           `json:"total_amount"`                                             // сумма
	Currency    string        `json:"currency"`                                                 // валюта "eur", "usd"
	Status      InvoiceStatus `json:"status" validate:"required,oneof=draft sent paid overdue"` //"draft", "sent", "paid", "overdue" (enum)
	Notes       *string       `json:"notes"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
type InvoiceItem struct {
	ID          uint   `json:"id"`
	InvoiceID   uint   `json:"invoice_id"`
	Description string `json:"description"` // описание товара, услуги
	Quantity    int    `json:"quantity"`
	UnitPrice   int    `json:"unit_price"`   //цена за единицу
	VatRate     int    `json:"vat_rate"`     // ставка ндс
	VatAmount   int    `json:"vat_amount"`   //сумма ндс
	TotalAmount int    `json:"total_amount"` //Quantity * UnitPrice + VatAmount
}
