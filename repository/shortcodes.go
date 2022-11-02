package repository

import (
	"context"

	"github.com/ochom/orctic-database/models"
)

// CreateSenderName ...
func (r *repo) CreateSenderName(ctx context.Context, data *models.SenderName) error {
	return r.DB.Create(data).Error
}

// UpdateSenderName ...
func (r *repo) UpdateSenderName(ctx context.Context, data *models.SenderName) error {
	return r.DB.Save(data).Error
}

// DeleteSenderName ...
func (r *repo) DeleteSenderName(ctx context.Context, query *models.SenderName) error {
	return r.DB.Where(query).Delete(&models.SenderName{}).Error
}

// GetSenderName ...
func (r *repo) GetSenderName(ctx context.Context, query *models.SenderName) (*models.SenderName, error) {
	var data models.SenderName
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetSenderNames ...
func (r *repo) GetSenderNames(ctx context.Context, query *models.SenderName) ([]*models.SenderName, error) {
	var data []*models.SenderName
	err := r.DB.Where(query).Find(&data).Error
	return data, err
}

// CreateOffer ...
func (r *repo) CreateOffer(ctx context.Context, data *models.Offer) error {
	return r.DB.Create(data).Error
}

// UpdateOffer ...
func (r *repo) UpdateOffer(ctx context.Context, data *models.Offer) error {
	return r.DB.Save(data).Error
}

// DeleteOffer ...
func (r *repo) DeleteOffer(ctx context.Context, query *models.Offer) error {
	return r.DB.Where(query).Delete(&models.Offer{}).Error
}

// GetOffer ...
func (r *repo) GetOffer(ctx context.Context, query *models.Offer) (*models.Offer, error) {
	var data models.Offer
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOffers ...
func (r *repo) GetOffers(ctx context.Context, query *models.Offer) ([]*models.Offer, error) {
	var data []*models.Offer
	err := r.DB.Where(query).Find(&data).Error
	return data, err
}
