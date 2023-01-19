package models

import (
	"sync"

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
