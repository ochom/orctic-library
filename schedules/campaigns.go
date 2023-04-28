package schedules

import (
	"time"

	"github.com/ochom/orctic-library/utils"
	"github.com/ochom/uuid"
	"gorm.io/gorm"
)

// Campaign ...
type Campaign struct {
	ID             string         `json:"id"`
	OrganizationID string         `json:"organizationID"`
	ContactGroupID string         `json:"contactGroupID"` // for messages targeted to a contact group
	SenderName     string         `json:"senderName"`     // this is the actual sender name/id
	OfferCode      string         `json:"offerCode"`      // this is the offer code
	OfferName      string         `json:"offerName"`      // this is the offer name
	OfferShortCode string         `json:"offerShortCode"` // this is the offer shortcode
	Type           CampaignType   `json:"type"`           // transactional or promotional
	Source         CampaignSource `json:"source"`
	Status         OutboxStatus   `json:"status"`
	Message        string         `json:"message"` // this is the template to be used for premium sms
	SendAt         time.Time      `json:"sendAt"`
	CreatedAt      time.Time      `json:"createdAt"`
	CreatedByID    string         `json:"createdByID"`
	IsPersonalized bool           `json:"isPersonalized"` // this tag determines if the messages will be sent in batches or individually
}

// NewCampaign ...
func NewCampaign(src CampaignSource, ct CampaignType, senderName, offerCode string, sendAt time.Time) *Campaign {
	return &Campaign{
		ID:         uuid.NewString(),
		SenderName: senderName,
		OfferCode:  offerCode,
		Type:       ct,
		SendAt:     sendAt,
		Source:     src,
		Status:     PendingOutbox,
	}
}

// Outbox ...
type Outbox struct {
	ID                string         `json:"id"`
	OrganizationID    string         `json:"organizationID"`
	CampaignID        string         `json:"campaignID"`
	LinkID            string         `json:"linkID"`     // for reply outbox
	SenderName        string         `json:"senderName"` // should be the actual sender name/id
	OfferCode         string         `json:"offerCode"`  // this is the offer code
	Recipient         string         `json:"recipient"`
	Message           string         `json:"message"`
	Cost              int            `json:"cost"` // number of units used for this outbox
	Status            OutboxStatus   `json:"status"`
	StatusDescription string         `json:"statusDescription"`
	CallbackURL       string         `json:"callbackURL"` // for outbox sent from API
	Retries           int            `json:"retries" gorm:"default:0"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `json:"deletedAt"`
}

// NewOutbox ...
func NewOutbox(campaignID, senderName, message, recipient string) *Outbox {
	return &Outbox{
		ID:                uuid.NewString(),
		CampaignID:        campaignID,
		SenderName:        senderName,
		Recipient:         recipient,
		Message:           message,
		Cost:              utils.GetMessageCost(message),
		Status:            PendingOutbox,
		StatusDescription: "Outbox is pending to be sent",
	}
}

// NewAPIOutbox ...
func NewAPIOutbox(campaignID, senderName, message, recipient, callbackURL string) *Outbox {
	outbox := NewOutbox(campaignID, senderName, message, recipient)
	outbox.CallbackURL = callbackURL
	return outbox
}
