package models

import (
	"time"

	"gorm.io/gorm"
)

// User ...
type User struct {
	ID         string         `json:"id" gorm:"primaryKey"`
	Username   string         `json:"username"`
	Mobile     string         `json:"mobile"`
	Email      string         `json:"email"`
	Password   string         `json:"-"`
	OTP        string         `json:"-"`
	IsVerified bool           `json:"isVerified"`
	IsAdmin    bool           `json:"isAdmin"`
	LastLogin  time.Time      `json:"lastLogin"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt"`
}
