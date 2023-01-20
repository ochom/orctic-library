package database

import "github.com/ochom/orctic-library/models"

// GetSchema return all models that can be migrated
func GetSchema() []any {
	return []any{
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
		&models.APIKey{},
		&models.Payment{},
		&models.Unit{},
		&models.ReplyOutbox{},
		&models.BulkOutbox{},
	}
}
