package models

import (
	"time"

	"gorm.io/gorm"
)

// Organization ...
type Organization struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	IsActive  bool           `json:"isActive"`
	CreatedBy string         `json:"createdBy"`
	UpdatedBy string         `json:"updatedBy"`
	DeletedBy string         `json:"deletedBy"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

// User ...
type User struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	OrganizationID string         `json:"organizationID"`
	Username       string         `json:"username"`
	Mobile         string         `json:"mobile"`
	Email          string         `json:"email"`
	Password       string         `json:"-"`
	OTP            string         `json:"-"`
	IsVerified     bool           `json:"isVerified"`
	IsAdmin        bool           `json:"isAdmin"`
	LastLogin      time.Time      `json:"lastLogin"`
	CreatedBy      string         `json:"createdBy"`
	UpdatedBy      string         `json:"updatedBy"`
	DeletedBy      string         `json:"deletedBy"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}

// ContactGroup ...
type ContactGroup struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	OrganizationID string         `json:"organizationID"`
	CreatedBy      string         `json:"createdBy"`
	UpdatedBy      string         `json:"updatedBy"`
	DeletedBy      string         `json:"deletedBy"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
	TotalContacts  int64          `json:"totalContacts" gorm:"-"`
}

// Contact ...
type Contact struct {
	ID        string         `json:"id"`
	GroupID   string         `json:"groupID"`
	Mobile    string         `json:"mobile"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
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
	OrganizationID string         `json:"organizationID"`
	Name           string         `json:"name" gorm:"uniqueIndex"`
	Description    string         `json:"description"`
	SenderNameType SenderNameType `json:"senderNameType"`
	CreatedBy      string         `json:"createdBy"`
	UpdatedBy      string         `json:"updatedBy"`
	DeletedBy      string         `json:"deletedBy"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}

// Offer ...
type Offer struct {
	ID             string         `json:"id"`
	OrganizationID string         `json:"organizationID"`
	Name           string         `json:"name" gorm:"uniqueIndex"`
	Description    string         `json:"description"`
	ShortCode      string         `json:"shortcode"`
	OfferCode      string         `json:"offerCode"`
	OfferType      OfferType      `json:"type"`
	CreatedBy      string         `json:"createdBy"`
	UpdatedBy      string         `json:"updatedBy"`
	DeletedBy      string         `json:"deletedBy"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
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
	CreatedBy      string              `json:"createdBy"`
	UpdatedBy      string              `json:"updatedBy"`
	DeletedBy      string              `json:"deletedBy"`
	CreatedAt      time.Time           `json:"createdAt"`
	UpdatedAt      time.Time           `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt      `json:"deletedAt"`
}

// Draft used for storing customized bulk campaigns
type Draft struct {
	ID             string         `json:"id"`
	OrganizationID string         `json:"organizationID"`
	CampaignID     string         `json:"campaignID"`
	Mobile         string         `json:"mobile"`
	Body           string         `json:"body"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}

// Outbox ...
type Outbox struct {
	ID                string         `json:"id"`
	OrganizationID    string         `json:"organizationID"`
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
