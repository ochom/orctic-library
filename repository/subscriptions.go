package repository

import (
	"context"

	"github.com/ochom/orctic-database/models"
)

func (r *repo) CreateSubscriber(ctx context.Context, data *models.Subscriber) error {
	return r.DB.Create(data).Error
}

func (r *repo) DeleteSubscriber(ctx context.Context, query *models.Subscriber) error {
	return r.DB.Where(query).Delete(&models.Subscriber{}).Error
}

func (r *repo) GetSubscriber(ctx context.Context, query *models.Subscriber) (*models.Subscriber, error) {
	var data models.Subscriber
	err := r.DB.Where(query).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *repo) GetSubscribers(ctx context.Context, query *models.Subscriber) ([]*models.Subscriber, error) {
	var data []*models.Subscriber
	err := r.DB.Where(query).Order("created_at desc").Find(&data).Error
	return data, err
}
