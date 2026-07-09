package models

import "time"

type ClientType string

const (
	ClientTypeIndividual ClientType = "individual"
	ClientTypeCompany    ClientType = "company"
)

type Client struct {
	ID   uint   `json:"id"`
	Type string `json:"type"` // "individual" или "company"
	// Общие поля
	Name  string `json:"name"` // для физлица - ФИО, для компании - название
	Phone string `json:"phone"`
	Email string `json:"email"`
	// Поля для компаний (могут быть NULL для физлиц)
	CompanyName        *string    `json:"company_name,omitempty"` // официальное название (если отличается от Name)
	RegistrationNumber *string    `json:"registration_number,omitempty"`
	VatNumber          *string    `json:"vat_number,omitempty"`
	ContactPerson      *string    `json:"contact_person,omitempty"`
	Addresses          []*Address `json:"address,omitempty"` // список адресов (загружается отдельно)
	// Мета-информация
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
