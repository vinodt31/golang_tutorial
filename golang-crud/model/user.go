package model

import "time"

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Model

	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
}
