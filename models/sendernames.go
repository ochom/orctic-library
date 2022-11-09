package models

import "gorm.io/gorm"

// SenderName ...
type SenderName struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"uniqueIndex"`
	Description    string         `json:"description"`
	SenderNameType SenderNameType `json:"senderNameType"`
	BaseModel
}

// AfterFind ...
func (s *SenderName) AfterFind(tx *gorm.DB) (err error) {
	if s.CreatedByID != "" {
		var createdBy User
		err = tx.Model(&User{}).Where("id = ?", s.CreatedByID).First(&createdBy).Error
		if err != nil {
			return err
		}
		s.CreatedBy = &createdBy
	}

	if s.UpdatedByID != "" {
		var updatedBy User
		err = tx.Model(&User{}).Where("id = ?", s.UpdatedByID).First(&updatedBy).Error
		if err != nil {
			return err
		}
		s.UpdatedBy = &updatedBy
	}

	return nil
}

// OrganizationSenderName ...
type OrganizationSenderName struct {
	OrganizationID string      `json:"organizationID"`
	SenderNameID   string      `json:"senderNameID"`
	SenderName     *SenderName `json:"senderName" gorm:"-"` // this is a virtual field
	BaseModel
}

// AfterFind ...
func (o *OrganizationSenderName) AfterFind(tx *gorm.DB) (err error) {
	if o.CreatedByID != "" {
		var createdBy User
		err = tx.Model(&User{}).Where("id = ?", o.CreatedByID).First(&createdBy).Error
		if err != nil {
			return err
		}
		o.CreatedBy = &createdBy
	}

	if o.UpdatedByID != "" {
		var updatedBy User
		err = tx.Model(&User{}).Where("id = ?", o.UpdatedByID).First(&updatedBy).Error
		if err != nil {
			return err
		}
		o.UpdatedBy = &updatedBy
	}

	var senderName SenderName
	err = tx.Model(&SenderName{}).Where("id = ?", o.SenderNameID).First(&senderName).Error
	if err != nil {
		return err
	}

	o.SenderName = &senderName

	return nil
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

// AfterFind ...
func (o *Offer) AfterFind(tx *gorm.DB) (err error) {
	if o.CreatedByID != "" {
		var createdBy User
		err = tx.Model(&User{}).Where("id = ?", o.CreatedByID).First(&createdBy).Error
		if err != nil {
			return err
		}
		o.CreatedBy = &createdBy
	}

	if o.UpdatedByID != "" {
		var updatedBy User
		err = tx.Model(&User{}).Where("id = ?", o.UpdatedByID).First(&updatedBy).Error
		if err != nil {
			return err
		}
		o.UpdatedBy = &updatedBy
	}

	return nil
}
