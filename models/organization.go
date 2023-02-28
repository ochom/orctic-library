package models

import (
	"fmt"

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
	Owner          *User  `json:"owner" gorm:"-"` // this is a virtual field
	BaseModel
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

	if err := tx.Model(&User{}).Where("id = ?", o.OwnerID).First(&o.Owner).Error; err != nil {
		return err
	}

	if err := tx.Model(&User{}).Where("id = ?", o.CreatedByID).Find(&o.CreatedBy).Error; err != nil {
		return err
	}

	return nil
}
