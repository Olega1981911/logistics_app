package models

import "time"

// Document — любой документ со сроком действия.
// Используется для комплаенса: страховки, техосмотры, лицензии, договоры.
type Document struct {
	ID       int64 `json:"id" db:"id"`
	TenantID int64 `json:"tenant_id" db:"tenant_id"`

	// К чему привязан: vehicle, driver, client, order, tenant
	EntityType string `json:"entity_type" db:"entity_type"`
	EntityID   int64  `json:"entity_id" db:"entity_id"`

	// Тип документа
	DocType string `json:"doc_type" db:"doc_type"` // insurance, inspection, license, contract, invoice

	// Номер документа (для поиска и сверки)
	Number string `json:"number" db:"number"`

	// Срок действия
	ValidFrom time.Time  `json:"valid_from" db:"valid_from"`
	ValidTo   *time.Time `json:"valid_to,omitempty" db:"valid_to"` // nil = бессрочный

	// Файл (ссылка на S3/локальное хранилище)
	FileURL string `json:"file_url" db:"file_url"`

	// Статус: valid, expiring_soon, expired
	// Обычно вычисляется в service по ValidTo, но кешируется для быстрого поиска
	Status string `json:"status" db:"status"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
