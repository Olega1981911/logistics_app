package models

import "time"

type UserRole string

const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleManager UserRole = "manager"
	UserRoleDriver  UserRole = "driver"
)

type User struct {
	ID           uint       `json:"id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"passwordhash"`                                        //bcrypt hash
	Role         UserRole   `json:"role" validate:"required,oneof=admin manager driver"` // "admin","manager","driver"
	Phone        *string    `json:"phone,omitempty"`
	Language     string     `json:"language"` //предпочитаемый язык
	Active       bool       `json:"active"`   //активный ли аккаунт
	LastLoginAt  *time.Time `json:"last_login_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
