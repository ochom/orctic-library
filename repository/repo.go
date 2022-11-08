package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/ochom/orctic-database/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:generate mockgen -source=database.go -destination=mocks/mock_database.go -package=mocks

// Platform ...
type Platform string

const (
	// MySQL ...
	MySQL Platform = "mysql"

	// Postgres ...
	Postgres Platform = "postgres"

	// SQLite ...
	SQLite Platform = "sqlite"
)

// Repo ...
type Repo interface {
	CreateOrganization(ctx context.Context, data *models.Organization) error
	UpdateOrganization(ctx context.Context, data *models.Organization) error
	DeleteOrganization(ctx context.Context, query *models.Organization) error
	GetOrganization(ctx context.Context, query *models.Organization) (*models.Organization, error)
	GetOrganizations(ctx context.Context, query *models.Organization) ([]*models.Organization, error)
	AddUserToOrganization(ctx context.Context, data *models.UserOrganization) error
	RemoveUserFromOrganization(ctx context.Context, data *models.UserOrganization) error
	GetUsersOrganizations(ctx context.Context, query *models.UserOrganization) ([]*models.Organization, error)
	GetOrganizationUsers(ctx context.Context, query *models.UserOrganization) ([]*models.User, error)

	CreateSenderName(ctx context.Context, data *models.SenderName) error
	UpdateSenderName(ctx context.Context, data *models.SenderName) error
	DeleteSenderName(ctx context.Context, query *models.SenderName) error
	GetSenderName(ctx context.Context, query *models.SenderName) (*models.SenderName, error)
	GetSenderNames(ctx context.Context, query *models.SenderName) ([]*models.SenderName, error)

	CreateOffer(ctx context.Context, data *models.Offer) error
	UpdateOffer(ctx context.Context, data *models.Offer) error
	DeleteOffer(ctx context.Context, query *models.Offer) error
	GetOffer(ctx context.Context, query *models.Offer) (*models.Offer, error)
	GetOffers(ctx context.Context, query *models.Offer) ([]*models.Offer, error)

	CreateUser(ctx context.Context, data *models.User) error
	UpdateUser(ctx context.Context, data *models.User) error
	DeleteUser(ctx context.Context, query *models.User) error
	GetUser(ctx context.Context, query *models.User) (*models.User, error)
	GetUsers(ctx context.Context, query *models.User) ([]*models.User, error)

	CreateSubscriber(ctx context.Context, data *models.Subscriber) error
	DeleteSubscriber(ctx context.Context, query *models.Subscriber) error
	GetSubscriber(ctx context.Context, query *models.Subscriber) (*models.Subscriber, error)
	GetSubscribers(ctx context.Context, query *models.Subscriber) ([]*models.Subscriber, error)

	CreateContactGroup(ctx context.Context, data *models.ContactGroup) error
	UpdateContactGroup(ctx context.Context, data *models.ContactGroup) error
	DeleteContactGroup(ctx context.Context, query *models.ContactGroup) error
	GetContactGroup(ctx context.Context, query *models.ContactGroup) (*models.ContactGroup, error)
	GetContactGroups(ctx context.Context, query *models.ContactGroup) ([]*models.ContactGroup, error)

	CreateContact(ctx context.Context, data *models.Contact) error
	UpdateContact(ctx context.Context, data *models.Contact) error
	DeleteContact(ctx context.Context, query *models.Contact) error
	GetContact(ctx context.Context, query *models.Contact) (*models.Contact, error)
	GetContacts(ctx context.Context, query *models.Contact) ([]*models.Contact, error)

	CreateCampaign(ctx context.Context, data *models.Campaign) error
	UpdateCampaign(ctx context.Context, data *models.Campaign) error
	DeleteCampaign(ctx context.Context, query *models.Campaign) error
	GetCampaign(ctx context.Context, query *models.Campaign) (*models.Campaign, error)
	GetCampaigns(ctx context.Context, query *models.Campaign) ([]*models.Campaign, error)

	CreateDraft(ctx context.Context, data *models.Draft) error
	UpdateDraft(ctx context.Context, data *models.Draft) error
	DeleteDraft(ctx context.Context, query *models.Draft) error
	GetDraft(ctx context.Context, query *models.Draft) (*models.Draft, error)
	GetDrafts(ctx context.Context, query *models.Draft) ([]*models.Draft, error)

	CreateOutbox(ctx context.Context, data *models.Outbox) error
	UpdateOutbox(ctx context.Context, data *models.Outbox) error
	DeleteOutbox(ctx context.Context, query *models.Outbox) error
	GetOutbox(ctx context.Context, query *models.Outbox) (*models.Outbox, error)
	GetOutboxes(ctx context.Context, query *models.Outbox) ([]*models.Outbox, error)
}

type repo struct {
	DB *gorm.DB
}

func (r *repo) init(pl Platform) error {
	dns, ok := os.LookupEnv("DATABASE_DNS")
	if !ok {
		return fmt.Errorf("DATABASE_DNS not set")
	}

	var db *gorm.DB
	var err error

	switch pl {
	case MySQL:
		db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	case Postgres:
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	case SQLite:
		db, err = gorm.Open(sqlite.Open(dns), &gorm.Config{})
	}

	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	r.DB = db

	schema := []interface{}{
		&models.Organization{},
		&models.User{},
		&models.UserOrganization{},
		&models.SenderName{},
		&models.OrganizationSenderName{},
		&models.Offer{},
		&models.OrganizationOffer{},
		&models.ContactGroup{},
		&models.Contact{},
		&models.Subscriber{},
		&models.Campaign{},
		&models.Draft{},
		&models.Outbox{},
	}

	err = r.DB.AutoMigrate(schema...)

	if err != nil {
		return fmt.Errorf("database migration failed: %s", err.Error())
	}

	return nil
}

// New ...
func New(dbPlatform Platform) (Repo, error) {
	r := &repo{}
	if err := r.init(dbPlatform); err != nil {
		return nil, err
	}

	return r, nil
}
