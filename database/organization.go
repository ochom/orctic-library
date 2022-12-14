package database

import (
	"context"

	"github.com/ochom/orctic-library/models"
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
func (r *repo) RemoveUserFromOrganization(ctx context.Context, query *models.UserOrganization) error {
	return r.DB.Where(query).Delete(&models.UserOrganization{}).Error
}

// GetUsersInOrganization ...
func (r *repo) GetUsersInOrganization(ctx context.Context, organizationID string) ([]*models.User, error) {
	var data []*models.User
	err := r.DB.Model(&models.User{}).Joins("JOIN user_organizations ON users.id = user_organizations.user_id").Where("user_organizations.organization_id = ?", organizationID).Find(&data).Error
	return data, err
}

// GetOrganizationsForUser ...
func (r *repo) GetOrganizationsForUser(ctx context.Context, userID string) ([]*models.Organization, error) {
	var data []*models.Organization
	err := r.DB.Model(&models.Organization{}).Joins("JOIN user_organizations ON organizations.id = user_organizations.organization_id").Where("user_organizations.user_id = ?", userID).Find(&data).Error
	return data, err
}

// CreateOrganizationKey ...
func (r *repo) CreateOrganizationKey(ctx context.Context, data *models.OrganizationKey) error {
	return r.DB.Create(data).Error
}

// UpdateOrganizationKey ...
func (r *repo) UpdateOrganizationKey(ctx context.Context, data *models.OrganizationKey) error {
	return r.DB.Save(data).Error
}

// DeleteOrganizationKey ...
func (r *repo) DeleteOrganizationKey(ctx context.Context, query *models.OrganizationKey) error {
	return r.DB.Where(query).Delete(&models.OrganizationKey{}).Error
}

// GetOrganizationKey ...
func (r *repo) GetOrganizationKey(ctx context.Context, query *models.OrganizationKey) (*models.OrganizationKey, error) {
	var data models.OrganizationKey
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOrganizationKeys ...
func (r *repo) GetOrganizationKeys(ctx context.Context, query *models.OrganizationKey) ([]*models.OrganizationKey, error) {
	var data []*models.OrganizationKey
	err := r.DB.Where(query).Order("created_at desc").Find(&data).Error
	return data, err
}
