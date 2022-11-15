package models

import (
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
	BaseModel
	SenderName *SenderName `json:"senderName" gorm:"-"` // this is a virtual field
	Offer      *Offer      `json:"offer" gorm:"-"`      // this is a virtual field
}

// AfterFind ...
func (c *Campaign) AfterFind(tx *gorm.DB) (err error) {
	if c.CreatedByID != "" {
		var createdBy User
		err = tx.Model(&User{}).Where("id = ?", c.CreatedByID).First(&createdBy).Error
		if err != nil {
			return err
		}
		c.CreatedBy = &createdBy
	}

	if c.UpdatedByID != "" {
		var updatedBy User
		err = tx.Model(&User{}).Where("id = ?", c.UpdatedByID).First(&updatedBy).Error
		if err != nil {
			return err
		}
		c.UpdatedBy = &updatedBy
	}

	if c.SenderNameID != "" {
		var senderName SenderName
		err = tx.Model(&SenderName{}).Where("id = ?", c.SenderNameID).First(&senderName).Error
		if err != nil {
			return err
		}
		c.SenderName = &senderName
	}

	if c.OfferID != "" {
		var offer Offer
		err = tx.Model(&Offer{}).Where("id = ?", c.OfferID).First(&offer).Error
		if err != nil {
			return err
		}
		c.Offer = &offer
	}

	return nil
}
