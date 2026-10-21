package models

import "time"

// AIAgentAction — одна операция ИИ-агента.
// Используется для аудита, отладки и обучения (что сработало, что нет).
type AIAgentAction struct {
	ID       int64 `json:"id" db:"id"`
	TenantID int64 `json:"tenant_id" db:"tenant_id"`

	// Какой агент: dispatcher, driver_assistant, accountant, economist, routing
	AgentType string `json:"agent_type" db:"agent_type"`

	// Что сделал: route_planned, invoice_audited, fuel_anomaly_detected, eta_recalculated
	ActionType string `json:"action_type" db:"action_type"`

	// К чему относится (если применимо)
	EntityType *string `json:"entity_type,omitempty" db:"entity_type"`
	EntityID   *int64  `json:"entity_id,omitempty" db:"entity_id"`

	// Вход и выход — JSON, потому что структура разная для каждого действия
	Input  string `json:"input" db:"input"`   // JSON
	Output string `json:"output" db:"output"` // JSON

	// Уверенность модели (0.0–1.0). Позволяет отсеивать «сомнительные» рекомендации.
	Confidence *float64 `json:"confidence,omitempty" db:"confidence"`

	// Принял ли человек рекомендацию: accepted, rejected, ignored, nil (ещё не решили)
	HumanDecision *string `json:"human_decision,omitempty" db:"human_decision"`

	// Модель, которая дала рекомендацию (для воспроизводимости)
	ModelUsed string `json:"model_used" db:"model_used"`

	// Сколько токенов/времени заняло (для аналитики стоимости ИИ)
	TokensUsed int `json:"tokens_used" db:"tokens_used"`
	LatencyMs  int `json:"latency_ms" db:"latency_ms"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
