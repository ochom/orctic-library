package models

import (
	"fmt"
	"sync"

	"github.com/ochom/orctic-library/utils"
	"gorm.io/gorm"
)

// Organization ...
type Organization struct {
	ID             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	IsActive       bool   `json:"isActive"`
	OwnerID        string `json:"ownerID" gorm:"required"`
	BillingAccount string `json:"billingAccount"`
	BaseModel
	Owner User `json:"owner" gorm:"-"` // this is a virtual field
}

// BeforeCreate ...
func (o *Organization) BeforeCreate(tx *gorm.DB) (err error) {
	maxAttempts := 10
	for {
		maxAttempts--
		if maxAttempts == 0 {
			return fmt.Errorf("unable to generate unique account number")
		}

		accountNumber := utils.GenerateOTP(5)
		var count int64
		err := tx.Model(&Organization{}).Where("billing_account = ?", accountNumber).Count(&count).Error
		if err != nil {
			return err
		}
		if count == 0 {
			o.BillingAccount = accountNumber
			break
		}
	}
	return nil
}

// AfterFind ...
func (o *Organization) AfterFind(tx *gorm.DB) (err error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	// set billing account if empty
	if o.BillingAccount == "" {
		maxAttempts := 10
		for {
			maxAttempts--
			if maxAttempts == 0 {
				return fmt.Errorf("unable to generate unique account number")
			}

			accountNumber := utils.GenerateOTP(5)
			var count int64
			err := tx.Model(&Organization{}).Where("billing_account = ?", accountNumber).Count(&count).Error
			if err != nil {
				return err
			}
			if count == 0 {
				o.BillingAccount = accountNumber
				_ = tx.Save(o)
				break
			}
		}
	}

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
