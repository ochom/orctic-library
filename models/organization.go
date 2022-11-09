package models

import (
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
	var owner User
	err = tx.Model(&User{}).Where("id = ?", o.OwnerID).First(&owner).Error
	if err != nil {
		return err
	}
	o.Owner = owner

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
