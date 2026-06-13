package models

import "time"

type User struct {
	ID           uint      `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	PhoneNumber  string    `json:"phone_number"`
	ProfileImage string    `json:"profile_image"`
	IsActive     bool      `json:"is_active"`
	Roles        []string  `json:"roles,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Address struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	AddressLine1 string    `json:"address_line1" binding:"required"`
	AddressLine2 string    `json:"address_line2"`
	Locality     string    `json:"locality"`
	City         string    `json:"city" binding:"required"`
	District     string    `json:"district"`
	State        string    `json:"state" binding:"required"`
	Country      string    `json:"country"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
