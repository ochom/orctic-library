package models

import (
	"sync"

	"gorm.io/gorm"
)

// SenderName ...
type SenderName struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"uniqueIndex"`
	Description    string         `json:"description"`
	PackageID      string         `json:"packageID"`
	SenderNameType SenderNameType `json:"senderNameType"`
	BaseModel
}

// AfterFind ...
func (s *SenderName) AfterFind(tx *gorm.DB) (err error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	// get created by
	wg.Add(1)
	go func() {
		defer wg.Done()
		var createdBy User
		err = tx.Model(&User{}).Where("id = ?", s.CreatedByID).First(&createdBy).Error
		if err != nil {
			return
		}

		mu.Lock()
		s.CreatedBy = &createdBy
		mu.Unlock()
	}()

	// get updated by
	wg.Add(1)
	go func() {
		defer wg.Done()
		var updatedBy User
		err = tx.Model(&User{}).Where("id = ?", s.UpdatedByID).First(&updatedBy).Error
		if err != nil {
			return
		}

		mu.Lock()
		s.UpdatedBy = &updatedBy
		mu.Unlock()
	}()
	wg.Wait()

	return nil
}

// OrganizationSenderName ...
type OrganizationSenderName struct {
	ID             string      `json:"id" gorm:"primaryKey"`
	OrganizationID string      `json:"organizationID"`
	SenderNameID   string      `json:"senderNameID"`
	CostPerSMS     float64     `json:"costPerSMS"`
	SenderName     *SenderName `json:"senderName" gorm:"-"` // this is a virtual field
	BaseModel
}

// AfterFind ...
func (o *OrganizationSenderName) AfterFind(tx *gorm.DB) (err error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	// get created by
	wg.Add(1)
	go func() {
		defer wg.Done()
		var createdBy User
		err = tx.Model(&User{}).Where("id = ?", o.CreatedByID).First(&createdBy).Error
		if err != nil {
			return
		}

		mu.Lock()
		o.CreatedBy = &createdBy
		mu.Unlock()
	}()

	// get updated by
	wg.Add(1)
	go func() {
		defer wg.Done()
		var updatedBy User
		err = tx.Model(&User{}).Where("id = ?", o.UpdatedByID).First(&updatedBy).Error
		if err != nil {
			return
		}

		mu.Lock()
		o.UpdatedBy = &updatedBy
		mu.Unlock()
	}()

	// get sender name
	wg.Add(1)
	go func() {
		defer wg.Done()
		var senderName SenderName
		err = tx.Model(&SenderName{}).Where("id = ?", o.SenderNameID).First(&senderName).Error
		if err != nil {
			return
		}

		mu.Lock()
		o.SenderName = &senderName
		mu.Unlock()
	}()

	wg.Wait()

	return nil
}

// Offer ...
type Offer struct {
	ID                 string        `json:"id"`
	OrganizationID     string        `json:"organizationID"`
	Name               string        `json:"name"`
	Description        string        `json:"description"`
	DisplayName        string        `json:"displayName"`
	DisplayDescription string        `json:"displayDescription"`
	ShowInWeb          bool          `json:"showInWeb"`
	DisplayIcon        string        `json:"displayIcon"`
	ShortCode          string        `json:"shortcode"`
	OfferCode          string        `json:"offerCode"`
	OfferType          OfferType     `json:"type"`
	ServerURL          string        `json:"serverURL"`                  // if the offer is an on-demand offer, this is the server url that will be called
	Organization       *Organization `json:"organization" gorm:"-"`      // this is a virtual field
	TotalSubscribers   int64         `json:"totalSubscribers" gorm:"-"`  // this is a virtual field
	ActiveSubscribers  int64         `json:"activeSubscribers" gorm:"-"` // this is a virtual field
	BaseModel
}

// AfterFind ...
func (o *Offer) AfterFind(tx *gorm.DB) (err error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	// get created by
	wg.Add(1)
	go func() {
		defer wg.Done()
		var createdBy User
		err = tx.Model(&User{}).Where("id = ?", o.CreatedByID).First(&createdBy).Error
		if err != nil {
			return
		}

		mu.Lock()
		o.CreatedBy = &createdBy
		mu.Unlock()
	}()

	// get updated by
	wg.Add(1)
	go func() {
		defer wg.Done()
		var updatedBy User
		err = tx.Model(&User{}).Where("id = ?", o.UpdatedByID).First(&updatedBy).Error
		if err != nil {
			return
		}

		mu.Lock()
		o.UpdatedBy = &updatedBy
		mu.Unlock()
	}()

	// get organization
	wg.Add(1)
	go func() {
		defer wg.Done()

		var organization Organization
		err = tx.Model(&Organization{}).Where("id = ?", o.OrganizationID).First(&organization).Error
		if err != nil {
			return
		}
		mu.Lock()
		o.Organization = &organization
		mu.Unlock()
	}()

	// get total subscribers
	wg.Add(1)
	go func() {
		defer wg.Done()

		var totalSubscribers int64
		var activeSubscribers int64
		err = tx.Model(&Subscriber{}).Where("offer_id = ?", o.ID).Count(&totalSubscribers).Error
		if err != nil {
			return
		}

		err = tx.Model(&Subscriber{}).Where("offer_id = ? AND status = ?", o.ID, ActiveSubscription).Count(&activeSubscribers).Error
		if err != nil {
			return
		}

		mu.Lock()
		o.TotalSubscribers = totalSubscribers
		o.ActiveSubscribers = activeSubscribers
		mu.Unlock()
	}()

	wg.Wait()

	return nil
}
