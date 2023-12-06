package schedules

import (
	"time"

	"github.com/ochom/uuid"
	"gorm.io/gorm"
)

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
	Source            CampaignSource `json:"source"`
	Status            OutboxStatus   `json:"status"`
	StatusDescription string         `json:"statusDescription"`
	ErrorDescription  string         `json:"errorDescription"`
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
		Cost:              getMessageCost(message),
		Status:            PendingOutbox,
		StatusDescription: "Outbox is pending to be sent",
	}
}

// getMessageCost ...
func getMessageCost(message string) int {
	pageSize := 160
	messageLength := len(message)

	if messageLength <= pageSize {
		return 1
	}

	if messageLength%pageSize == 0 {
		return messageLength / pageSize
	}

	return (messageLength / pageSize) + 1
}

// Inbox CP notification inbox
type Inbox struct {
	ID        string         `json:"id"`
	RequestID string         `json:"requestID"`
	OfferCode string         `json:"offerCode"`
	LinkID    string         `json:"linkID"`
	Mobile    string         `json:"mobile"`
	Message   string         `json:"message"`
	Response  string         `json:"response"`
	Replied   bool           `json:"replied"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
