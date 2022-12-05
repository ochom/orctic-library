package models

import (
	"sync"
	"time"

	"gorm.io/gorm"
)

// Campaign ...
type Campaign struct {
	ID             string              `json:"id"`
	OrganizationID string              `json:"organizationID"`
	SenderNameID   string              `json:"senderNameID"`
	OfferID        string              `json:"offerID"`
	CampaignType   CampaignType        `json:"campaignType"`
	MessageType    CampaignMessageType `json:"messageType"`
	Template       string              `json:"template"`
	ContactGroupID string              `json:"contactGroupID"`
	SendAt         time.Time           `json:"sendAt"`
	Status         CampaignStatus      `json:"status"`
	BaseModel
	SenderName *SenderName `json:"senderName" gorm:"-"` // this is a virtual field
	Offer      *Offer      `json:"offer" gorm:"-"`      // this is a virtual field
}

// AfterFind ...
func (c *Campaign) AfterFind(tx *gorm.DB) (err error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	// get sender name
	wg.Add(1)
	go func() {
		defer wg.Done()
		var senderName SenderName
		err = tx.Model(&SenderName{}).Where("id = ?", c.SenderNameID).First(&senderName).Error
		if err != nil {
			return
		}

		mu.Lock()
		c.SenderName = &senderName
		mu.Unlock()
	}()

	// get offer
	wg.Add(1)
	go func() {
		defer wg.Done()
		var offer Offer
		err = tx.Model(&Offer{}).Where("id = ?", c.OfferID).First(&offer).Error
		if err != nil {
			return
		}
		mu.Lock()
		c.Offer = &offer
		mu.Unlock()
	}()

	// get created by
	wg.Add(1)
	go func() {
		defer wg.Done()
		var createdBy User
		err = tx.Model(&User{}).Where("id = ?", c.CreatedByID).First(&createdBy).Error
		if err != nil {
			return
		}

		mu.Lock()
		c.CreatedBy = &createdBy
		mu.Unlock()
	}()

	// get updated by
	wg.Add(1)
	go func() {
		defer wg.Done()
		var updatedBy User
		err = tx.Model(&User{}).Where("id = ?", c.UpdatedByID).First(&updatedBy).Error
		if err != nil {
			return
		}

		mu.Lock()
		c.UpdatedBy = &updatedBy
		mu.Unlock()
	}()
	wg.Wait()

	return nil
}
