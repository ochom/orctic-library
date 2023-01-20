package database

import (
	grm "github.com/ochom/generic-gorm"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Repository ...
type Repository struct {
	*gorm.DB
}

// NewRepository ...
func NewRepository(pl grm.Platform, dsn string, level logger.LogLevel) (*Repository, error) {
	sql := grm.Init(pl, dsn, level)
	err := grm.Migrate(getSchema()...)
	if err != nil {
		return nil, err
	}

	return &Repository{
		DB: sql.DB,
	}, nil
}
