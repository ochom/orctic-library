package models

import (
	"sync"

	"gorm.io/gorm"
)

// Organization ...
type Organization struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	IsActive bool   `json:"isActive"`
	OwnerID  string `json:"ownerID" gorm:"required"`
	BaseModel
	Owner User `json:"owner" gorm:"-"` // this is a virtual field
}

// AfterFind ...
func (o *Organization) AfterFind(tx *gorm.DB) (err error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	// get owner
	wg.Add(1)
	go func() {
		defer wg.Done()
		var owner User
		err = tx.Model(&User{}).Where("id = ?", o.OwnerID).First(&owner).Error
		if err != nil {
			return
		}
		mu.Lock()
		o.Owner = owner
		mu.Unlock()
	}()

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
	wg.Wait()
	return nil
}

// OrganizationKey ...
type OrganizationKey struct {
	ID             string        `json:"id" gorm:"primaryKey"`
	OrganizationID string        `json:"organizationID" gorm:"required"`
	Key            string        `json:"key" gorm:"required"`
	Organization   *Organization `json:"organization" gorm:"-"` // this is a virtual field
	BaseModel
}

// AfterFind ...
func (ok *OrganizationKey) AfterFind(tx *gorm.DB) (err error) {
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
