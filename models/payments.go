package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Payment ...
type Payment struct {
	ID                string        `json:"id"`
	OrganizationID    string        `json:"organizationID"`
	Source            PaymentSource `json:"source"`
	Mobile            string        `json:"mobile"`
	Amount            string        `json:"amount"`
	Status            PaymentStatus `json:"status"`
	StatusDescription string        `json:"statusDescription"`
	ReferenceID       string        `json:"referenceID"`
	BaseModel
}

// AfterFind ...
func (p *Payment) AfterFind(tx *gorm.DB) error {
	var user User
	if err := tx.Where("id = ?", p.CreatedByID).First(&user).Error; err != nil {
		return fmt.Errorf("created by id not found: %v", err)
	}

	p.CreatedBy = &user
	return nil
}

// Unit ...
type Unit struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationID"`
	PaymentID      string `json:"paymentID"`
	Allocated      int    `json:"allocated"`
	Used           int    `json:"used"`
	BaseModel
}

// AfterFind ...
func (u *Unit) AfterFind(tx *gorm.DB) error {
	var user User
	if err := tx.Where("id = ?", u.CreatedByID).First(&user).Error; err != nil {
		return fmt.Errorf("created by id not found: %v", err)
	}

	u.CreatedBy = &user
	return nil
}
