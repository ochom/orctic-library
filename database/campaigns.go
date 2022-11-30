package database

import (
	"context"

	"github.com/ochom/orctic-library/models"
)

// CreateCampaign ...
func (r *repo) CreateCampaign(ctx context.Context, data *models.Campaign) error {
	return r.DB.Create(data).Error
}

// UpdateCampaign ...
func (r *repo) UpdateCampaign(ctx context.Context, data *models.Campaign) error {
	return r.DB.Save(data).Error
}

// DeleteCampaign ...
func (r *repo) DeleteCampaign(ctx context.Context, query *models.Campaign) error {
	return r.DB.Where(query).Delete(&models.Campaign{}).Error
}

// GetCampaign ...
func (r *repo) GetCampaign(ctx context.Context, query *models.Campaign) (*models.Campaign, error) {
	var data models.Campaign
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetCampaigns ...
func (r *repo) GetCampaigns(ctx context.Context, query *models.Campaign) ([]*models.Campaign, error) {
	var data []*models.Campaign
	err := r.DB.Where(query).Order("created_at desc").Find(&data).Error
	return data, err
}

// CreateDraft ...
func (r *repo) CreateDraft(ctx context.Context, data *models.Draft) error {
	return r.DB.Create(data).Error
}

// UpdateDraft ...
func (r *repo) UpdateDraft(ctx context.Context, data *models.Draft) error {
	return r.DB.Save(data).Error
}

// DeleteDraft ...
func (r *repo) DeleteDraft(ctx context.Context, query *models.Draft) error {
	return r.DB.Where(query).Delete(&models.Draft{}).Error
}

// GetDraft ...
func (r *repo) GetDraft(ctx context.Context, query *models.Draft) (*models.Draft, error) {
	var data models.Draft
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDrafts ...
func (r *repo) GetDrafts(ctx context.Context, query *models.Draft) ([]*models.Draft, error) {
	var data []*models.Draft
	err := r.DB.Where(query).Order("created_at desc").Find(&data).Error
	return data, err
}

// CreateOutbox ...
func (r *repo) CreateOutbox(ctx context.Context, data *models.Outbox) error {
	return r.DB.Create(data).Error
}

// UpdateOutbox ...
func (r *repo) UpdateOutbox(ctx context.Context, data *models.Outbox) error {
	return r.DB.Save(data).Error
}

// DeleteOutbox ...
func (r *repo) DeleteOutbox(ctx context.Context, query *models.Outbox) error {
	return r.DB.Where(query).Delete(&models.Outbox{}).Error
}

// GetOutbox ...
func (r *repo) GetOutbox(ctx context.Context, query *models.Outbox) (*models.Outbox, error) {
	var data models.Outbox
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOutboxes ...
func (r *repo) GetOutboxes(ctx context.Context, query *models.Outbox) ([]*models.Outbox, error) {
	var data []*models.Outbox
	err := r.DB.Where(query).Order("created_at desc").Find(&data).Error
	return data, err
}
