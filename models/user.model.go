package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	USER_ROLE_USER  UserRole = "user"
	USER_ROLE_ADMIN UserRole = "admin"
)

// User the user object model
type User struct {
	Id           uuid.UUID `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	LastName     string    `json:"last_name"`
	FirstName    string    `json:"first_name"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// SignUpDto the sign up data transfer object
type SignUpDto struct {
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,is_password"`
	LastName  string    `json:"lastName" validate:"required,min=2,max=25"`
	FirstName string    `json:"firstName" validate:"required,min=2,max=25"`
	Role      *UserRole `json:"role" validate:"omitempty,is_enum"`
}

// SignInDto the sign in data transfer object
type SignInDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,is_password"`
}

// AuthenticatedUser the authenticated user object
type AuthenticatedUser struct {
	User        *User  `json:"user"`
	AccessToken string `json:"accessToken"`
}

// IsValid checks if status is valid
func (u UserRole) IsValid() bool {
	switch u {
	case USER_ROLE_ADMIN, USER_ROLE_USER:
		return true
	}
	return false
}
