package database

import (
	grm "github.com/ochom/generic-gorm"
	"gorm.io/gorm"
)

// Repo ...
type Repo struct {
	DB *gorm.DB
}

// NewRepository ...
func NewRepository(pl grm.Platform, dsn string) (*Repo, error) {
	db := grm.Init(pl, dsn)
	if err := grm.Migrate(getSchema()...); err != nil {
		return nil, err
	}

	return &Repo{
		DB: db.DB,
	}, nil
}
