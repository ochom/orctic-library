package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel ...
type BaseModel struct {
	CreatedByID string `json:"createdByID"`
	UpdatedByID string `json:"updatedByID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	CreatedBy   User           `json:"createdBy" gorm:"-"` // this is a virtual field
	UpdatedBy   User           `json:"updatedBy" gorm:"-"` // this is a virtual field
}

// Organization ...
type Organization struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	IsActive bool   `json:"isActive"`
	OwnerID  string `json:"ownerID"`
	BaseModel
	Owner User `json:"owner" gorm:"-"` // this is a virtual field
}

// User ...
type User struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	OrganizationID string    `json:"organizationID"`
	Username       string    `json:"username"`
	Mobile         string    `json:"mobile"`
	Email          string    `json:"email"`
	Password       string    `json:"-"`
	OTP            string    `json:"-"`
	IsVerified     bool      `json:"isVerified"`
	IsAdmin        bool      `json:"isAdmin"`
	LastLogin      time.Time `json:"lastLogin"`
}

// UserOrganization ...
type UserOrganization struct {
	UserID         string `json:"userID"`
	OrganizationID string `json:"organizationID"`
	IsOwner        bool   `json:"isOwner"`
}

// ContactGroup ...
type ContactGroup struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organizationID"`
	BaseModel
}

// Contact ...
type Contact struct {
	ID      string `json:"id"`
	GroupID string `json:"groupID"`
	Mobile  string `json:"mobile"`
	BaseModel
}

// Subscriber ...
type Subscriber struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Mobile    string         `json:"mobile"`
	OfferID   string         `json:"offerID"`
	IsActive  bool           `json:"isActive"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

// SenderName ...
type SenderName struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"uniqueIndex"`
	Description    string         `json:"description"`
	SenderNameType SenderNameType `json:"senderNameType"`
	BaseModel
}

// OrganizationSenderName ...
type OrganizationSenderName struct {
	ID             string `json:"id" gorm:"primaryKey"`
	OrganizationID string `json:"organizationID"`
	SenderNameID   string `json:"senderNameID"`
}

// Offer ...
type Offer struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organizationID"`
	Name           string    `json:"name" gorm:"uniqueIndex"`
	Description    string    `json:"description"`
	ShortCode      string    `json:"shortcode"`
	OfferCode      string    `json:"offerCode"`
	OfferType      OfferType `json:"type"`
	BaseModel
}

// OrganizationOffer ...
type OrganizationOffer struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationID"`
	OfferID        string `json:"offerID"`
}

// Campaign ...
type Campaign struct {
	ID             string              `json:"id"`
	OrganizationID string              `json:"organizationID"`
	SenderNameID   string              `json:"senderNameID"`
	SenderName     string              `json:"senderName"`
	OfferID        string              `json:"offerID"`
	OfferCode      string              `json:"offerCode"`
	CampaignType   CampaignType        `json:"campaignType"`
	MessageType    CampaignMessageType `json:"messageType"`
	Template       string              `json:"template"`
	ContactGroupID string              `json:"contactGroupID"`
	SendAt         time.Time           `json:"sendAt"`
	BaseModel
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
	CampaignID        string         `json:"campaignID"`
	Mobile            string         `json:"mobile"`
	Body              string         `json:"body"`
	LinkID            string         `json:"linkID"`
	Status            string         `json:"status"`
	StatusDescription string         `json:"statusDescription"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"deletedAt"`
}
