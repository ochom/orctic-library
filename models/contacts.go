package models

import (
	"sync"

	"gorm.io/gorm"
)

// ContactGroup ...
type ContactGroup struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organizationID"`
	BaseModel
}

// AfterFind ...
func (c *ContactGroup) AfterFind(tx *gorm.DB) (err error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

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

// Contact ...
type Contact struct {
	ID      string `json:"id"`
	GroupID string `json:"groupID"`
	Mobile  string `json:"mobile"`
	BaseModel
}

// AfterFind ...
func (c *Contact) AfterFind(tx *gorm.DB) (err error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

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
