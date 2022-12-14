package database

import (
	"context"
	"fmt"
	"os"

	"github.com/ochom/orctic-library/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

	CreateOrganizationKey(ctx context.Context, data *models.OrganizationKey) error
	UpdateOrganizationKey(ctx context.Context, data *models.OrganizationKey) error
	DeleteOrganizationKey(ctx context.Context, query *models.OrganizationKey) error
	GetOrganizationKey(ctx context.Context, query *models.OrganizationKey) (*models.OrganizationKey, error)
	GetOrganizationKeys(ctx context.Context, query *models.OrganizationKey) ([]*models.OrganizationKey, error)

	CreateSenderName(ctx context.Context, data *models.SenderName) error
	UpdateSenderName(ctx context.Context, data *models.SenderName) error
	DeleteSenderName(ctx context.Context, query *models.SenderName) error
	GetSenderName(ctx context.Context, query *models.SenderName) (*models.SenderName, error)
	GetSenderNames(ctx context.Context, query *models.SenderName) ([]*models.SenderName, error)

	AddSenderNameToOrganization(ctx context.Context, data *models.OrganizationSenderName) error
	UpdateOrganizationSenderName(ctx context.Context, data *models.OrganizationSenderName) error
	RemoveSenderNameFromOrganization(ctx context.Context, query *models.OrganizationSenderName) error
	GetOrganizationSenderNames(ctx context.Context, orgID string) ([]*models.OrganizationSenderName, error)

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

	AddUserToOrganization(ctx context.Context, data *models.UserOrganization) error
	RemoveUserFromOrganization(ctx context.Context, query *models.UserOrganization) error
	GetUsersInOrganization(ctx context.Context, organizationID string) ([]*models.User, error)
	GetOrganizationsForUser(ctx context.Context, userID string) ([]*models.Organization, error)

	CreateSubscriber(ctx context.Context, data *models.Subscriber) error
	UpdateSubscriber(ctx context.Context, data *models.Subscriber) error
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

	sqlLogger := logger.Default.LogMode(logger.Silent)

	switch pl {
	case MySQL:
		db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
			Logger: sqlLogger,
		})
	case Postgres:
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{
			Logger: sqlLogger,
		})
	case SQLite:
		db, err = gorm.Open(sqlite.Open(dns), &gorm.Config{
			Logger: sqlLogger,
		})
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
		&models.ContactGroup{},
		&models.Contact{},
		&models.Subscriber{},
		&models.Campaign{},
		&models.Draft{},
		&models.Outbox{},
		&models.OrganizationKey{},
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
