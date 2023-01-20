package models

import (
	"strings"
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// APIKey ...
type APIKey struct {
	ID             string        `json:"id" gorm:"primaryKey"`
	OrganizationID string        `json:"organizationID"`
	Name           string        `json:"name"`
	Token          string        `json:"token"`
	Organization   *Organization `json:"organization" gorm:"-"` // this is a virtual field
	BaseModel
}

// AfterFind ...
func (ok *APIKey) AfterFind(tx *gorm.DB) (err error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	// get organization
	wg.Add(1)
	go func() {
		defer wg.Done()
		var organization Organization
		err = tx.Model(&Organization{}).Where("id = ?", ok.OrganizationID).First(&organization).Error
		if err != nil {
			return
		}

		mu.Lock()
		ok.Organization = &organization
		mu.Unlock()
	}()

	// get created by
	wg.Add(1)
	go func() {
		defer wg.Done()
		var createdBy User
		err = tx.Model(&User{}).Where("id = ?", ok.CreatedByID).First(&createdBy).Error
		if err != nil {
			return
		}

		mu.Lock()
		ok.CreatedBy = &createdBy
		mu.Unlock()
	}()

	// get updated by
	wg.Add(1)
	go func() {
		defer wg.Done()
		var updatedBy User
		err = tx.Model(&User{}).Where("id = ?", ok.UpdatedByID).First(&updatedBy).Error
		if err != nil {
			return
		}

		mu.Lock()
		ok.UpdatedBy = &updatedBy
		mu.Unlock()
	}()

	wg.Wait()
	return nil
}

// ReplyOutbox this is the API outbox for offerCSode replies
type ReplyOutbox struct {
	ID                string       `json:"id" gorm:"primaryKey"`
	LinkID            string       `json:"linkID"`
	OfferCode         string       `json:"offerCode"`
	Mobile            string       `json:"mobile"`
	Message           string       `json:"message"`
	Status            OutboxStatus `json:"status"`
	StatusDescription string       `json:"statusDescription"`
	BaseModel
}

// BulkOutbox this is the API outbox for broadcast messages
type BulkOutbox struct {
	ID                string              `json:"id" gorm:"primaryKey"`
	Type              CampaignMessageType `json:"type"`
	SenderName        string              `json:"senderName"`
	Message           string              `json:"message"`
	Recipients        string              `json:"recipients"`
	CallbackURL       string              `json:"callbackURL"`
	Units             int                 `json:"units"`
	Status            OutboxStatus        `json:"status"`
	StatusDescription string              `json:"statusDescription"`
	BaseModel
}

// GetRecipients ...
func (bo *BulkOutbox) GetRecipients() []string {
	switch bo.Type {
	case BroadcastCampaignMessage:
		numbers := []string{}
		for _, number := range strings.Split(bo.Recipients, ",") {
			numbers = append(numbers, strings.TrimSpace(number))
		}
		return numbers

	default:
		return []string{bo.Recipients}
	}
}

// NewBroadcastOutbox ...
func NewBroadcastOutbox(senderName, message string, recipients []string, callbackURL string, units int) *BulkOutbox {
	return &BulkOutbox{
		ID:          uuid.NewString(),
		Type:        BroadcastCampaignMessage,
		SenderName:  senderName,
		Message:     message,
		Recipients:  strings.Join(recipients, ","),
		CallbackURL: callbackURL,
		Units:       units,
	}
}

// NewPersonalizedOutbox ...
func NewPersonalizedOutbox(senderName, message, recipient, callbackURL string, units int) *BulkOutbox {
	return &BulkOutbox{
		ID:          uuid.NewString(),
		Type:        PersonalizedCampaignMessage,
		SenderName:  senderName,
		Message:     message,
		Recipients:  recipient,
		CallbackURL: callbackURL,
		Units:       units,
	}
}
