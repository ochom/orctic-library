package models

import "gorm.io/gorm"

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

	return nil
}

//
// Contact ...
type Contact struct {
	ID      string `json:"id"`
	GroupID string `json:"groupID"`
	Mobile  string `json:"mobile"`
	BaseModel
}

// AfterFind ...
func (c *Contact) AfterFind(tx *gorm.DB) (err error) {
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

	return nil
}
