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

// UserOrganization ...
type UserOrganization struct {
	UserID         string `json:"userID"`
	OrganizationID string `json:"organizationID"`
}

// Subscriber ...
type Subscriber struct {
	ID                string             `json:"id" gorm:"primaryKey"`
	Mobile            string             `json:"mobile"`
	OfferID           string             `json:"offerID"`
	IsActive          bool               `json:"isActive"`
	Source            SubscriptionSource `json:"source"`
	Status            SubscriberStatus   `json:"status"`
	StatusDescription string             `json:"statusDescription"`
	CreatedAt         time.Time          `json:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt     `json:"deletedAt"`
}

// Draft used for storing customized bulk campaigns
type Draft struct {
	ID         string         `json:"id"`
	CampaignID string         `json:"campaignID"`
	Mobile     string         `json:"mobile"`
	Body       string         `json:"body"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt"`
}

// Outbox ...
type Outbox struct {
	ID                string         `json:"id"`
	OrganizationID    string         `json:"organizationID"`
	CampaignID        string         `json:"campaignID"`
	Mobile            string         `json:"mobile"`
	Body              string         `json:"body"`
	LinkID            string         `json:"linkID"`
	Source            OutboxSource   `json:"source"`
	Status            OutboxStatus   `json:"status"`
	StatusDescription string         `json:"statusDescription"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"deletedAt"`
}
