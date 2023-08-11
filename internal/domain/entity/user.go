package entity

import "time"

type User struct {
	ID       int
	FullName string `json:"full_name"`
	Email    string `json:"email"`

	Phone    *string `json:"phone"`
	Password *string `json:"-"`
	Role     string  `json:"role"`
	IsActive bool    `json:"-"`

	EmailConfirmedAt  *time.Time `json:"email_confirm_at"`
	PhoneConfirmedAt  *time.Time `json:"phone_confirm_at"`
	PasswordChangedAt *time.Time `json:"password_changed_at"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	LastLoginAt   *time.Time `json:"last_login_at"`
	LastUserAgent *string    `json:"last_user_agent"`
	LastUserIP    *string    `json:"last_user_ip"`
}

type UserFilter struct {
	FullName string
	Phone    string
	Password string
	Role     string
	IsActive bool
	Email    string
	ID       int
}
