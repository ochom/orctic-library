package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel ...
type BaseModel struct {
	CreatedByID string `json:"createdByID"  gorm:"required"`
	UpdatedByID string `json:"updatedByID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	CreatedBy   *User          `json:"createdBy" gorm:"-"` // this is a virtual field
	UpdatedBy   *User          `json:"updatedBy" gorm:"-"` // this is a virtual field
}
