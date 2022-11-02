package repository

import (
	"context"

	"github.com/ochom/orctic-database/models"
)

// CreateContactGroup ...
func (r *repo) CreateContactGroup(ctx context.Context, data *models.ContactGroup) error {
	return r.DB.Create(data).Error
}

// UpdateContactGroup ...
func (r *repo) UpdateContactGroup(ctx context.Context, data *models.ContactGroup) error {
	return r.DB.Save(data).Error
}

// DeleteContactGroup ...
func (r *repo) DeleteContactGroup(ctx context.Context, query *models.ContactGroup) error {
	return r.DB.Where(query).Delete(&models.ContactGroup{}).Error
}

// GetContactGroup ...
func (r *repo) GetContactGroup(ctx context.Context, query *models.ContactGroup) (*models.ContactGroup, error) {
	var data models.ContactGroup
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetContactGroups ...
func (r *repo) GetContactGroups(ctx context.Context, query *models.ContactGroup) ([]*models.ContactGroup, error) {
	var data []*models.ContactGroup
	err := r.DB.Where(query).Find(&data).Error
	return data, err
}

// CreateContact ...
func (r *repo) CreateContact(ctx context.Context, data *models.Contact) error {
	return r.DB.Create(data).Error
}

// UpdateContact ...
func (r *repo) UpdateContact(ctx context.Context, data *models.Contact) error {
	return r.DB.Save(data).Error
}

// DeleteContact ...
func (r *repo) DeleteContact(ctx context.Context, query *models.Contact) error {
	return r.DB.Where(query).Delete(&models.Contact{}).Error
}

// GetContact ...
func (r *repo) GetContact(ctx context.Context, query *models.Contact) (*models.Contact, error) {
	var data models.Contact
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetContacts ...
func (r *repo) GetContacts(ctx context.Context, query *models.Contact) ([]*models.Contact, error) {
	var data []*models.Contact
	err := r.DB.Where(query).Find(&data).Error
	return data, err
}
