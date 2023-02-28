package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Campaign ...
type Campaign struct {
	ID             string          `json:"id"`
	OrganizationID string          `json:"organizationID"`
	ContactGroupID string          `json:"contactGroupID"` // for messages targeted to a contact group
	SenderName     string          `json:"senderName"`     // this is the actual sender name/id
	OfferCode      string          `json:"offerCode"`      // this is the offer code
	Channel        CampaignChannel `json:"channel"`        // bulk or premium
	Scheme         CampaignScheme  `json:"scheme"`         // personalized or broadcast
	Type           CampaignType    `json:"type"`           // transactional or promotional
	Source         CampaignSource  `json:"source"`
	Status         CampaignStatus  `json:"status"`
	Template       string          `json:"template"`
	SendAt         time.Time       `json:"sendAt"`
	BaseModel
}

// NewCampaign ...
func NewCampaign(src CampaignSource, ch CampaignChannel, ct CampaignType, sc CampaignScheme, senderName, offerCode, template, userID string, sendAt time.Time) *Campaign {
	return &Campaign{
		ID:         uuid.NewString(),
		SenderName: senderName,
		OfferCode:  offerCode,
		Channel:    ch,
		Scheme:     sc,
		Type:       ct,
		Template:   template,
		SendAt:     sendAt,
		Source:     src,
		Status:     PendingCampaign,
		BaseModel:  BaseModel{CreatedByID: userID},
	}
}

// AfterFind ...
func (c *Campaign) AfterFind(tx *gorm.DB) error {
	if err := tx.Model(&User{}).Where("id = ?", c.CreatedByID).First(&c.CreatedBy).Error; err != nil {
		return err
	}
	return nil
}

// Outbox ...
type Outbox struct {
	ID                string         `json:"id"`
	OrganizationID    string         `json:"organizationID"`
	CampaignID        string         `json:"campaignID"`
	LinkID            string         `json:"linkID"` // for reply outbox
	BatchID           string         `json:"batchID"`
	SenderName        string         `json:"senderName"` // should be the actual sender name/id
	Recipient         string         `json:"recipient"`
	Message           string         `json:"message"`
	Units             int            `json:"units"` // number of units used for this outbox
	Status            OutboxStatus   `json:"status"`
	StatusDescription string         `json:"statusDescription"`
	CallbackURL       string         `json:"callbackURL"` // for outbox sent from API
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"deletedAt"`
}

// NewOutbox ...
func NewOutbox(units int, campaignID, senderName, message, recipient string) *Outbox {
	return &Outbox{
		ID:                uuid.NewString(),
		CampaignID:        campaignID,
		SenderName:        senderName,
		Recipient:         recipient,
		Message:           message,
		Units:             units,
		Status:            PendingOutbox,
		StatusDescription: "Outbox is pending to be sent",
	}
}

// NewAPIOutbox ...
func NewAPIOutbox(units int, campaignID, senderName, message, recipient, callbackURL string) *Outbox {
	outbox := NewOutbox(units, campaignID, senderName, message, recipient)
	outbox.CallbackURL = callbackURL
	return outbox
}
