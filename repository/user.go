package repository

import (
	"context"

	"github.com/ochom/orctic-database/models"
)

// CreateUser ...
func (r *repo) CreateUser(ctx context.Context, data *models.User) error {
	return r.DB.Create(data).Error
}

// UpdateUser ...
func (r *repo) UpdateUser(ctx context.Context, data *models.User) error {
	return r.DB.Save(data).Error
}

// DeleteUser ...
func (r *repo) DeleteUser(ctx context.Context, query *models.User) error {
	return r.DB.Where(query).Delete(&models.User{}).Error
}

// GetUser ...
func (r *repo) GetUser(ctx context.Context, query *models.User) (*models.User, error) {
	var data models.User
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetUsers ...
func (r *repo) GetUsers(ctx context.Context, query *models.User) ([]*models.User, error) {
	var data []*models.User
	err := r.DB.Where(query).Order("created_at desc").Find(&data).Error
	return data, err
}
