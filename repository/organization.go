package repository

import (
	"context"

	"github.com/ochom/orctic-database/models"
)

// CreateOrganization ...
func (r *repo) CreateOrganization(ctx context.Context, data *models.Organization) error {
	return r.DB.Create(data).Error
}

// UpdateOrganization ...
func (r *repo) UpdateOrganization(ctx context.Context, data *models.Organization) error {
	return r.DB.Save(data).Error
}

// DeleteOrganization ...
func (r *repo) DeleteOrganization(ctx context.Context, query *models.Organization) error {
	return r.DB.Where(query).Delete(&models.Organization{}).Error
}

// GetOrganization ...
func (r *repo) GetOrganization(ctx context.Context, query *models.Organization) (*models.Organization, error) {
	var data models.Organization
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOrganizations ...
func (r *repo) GetOrganizations(ctx context.Context, query *models.Organization) ([]*models.Organization, error) {
	var data []*models.Organization
	err := r.DB.Where(query).Order("created_at desc").Find(&data).Error
	return data, err
}

// AddUserToOrganization ...
func (r *repo) AddUserToOrganization(ctx context.Context, data *models.UserOrganization) error {
	return r.DB.Create(data).Error
}

// RemoveUserFromOrganization ...
func (r *repo) RemoveUserFromOrganization(ctx context.Context, data *models.UserOrganization) error {
	return r.DB.Delete(data).Error
}

// GetOrganizationUsers ...
func (r *repo) GetOrganizationUsers(ctx context.Context, query *models.UserOrganization) ([]*models.User, error) {
	var data []*models.User
	err := r.DB.Where(query).Find(&data).Error
	return data, err
}

// GetUsersOrganizations ...
func (r *repo) GetUsersOrganizations(ctx context.Context, query *models.UserOrganization) ([]*models.Organization, error) {
	var data []*models.Organization
	err := r.DB.Where(query).Find(&data).Error
	return data, err
}
