package schedules

import (
	"time"

	"github.com/ochom/uuid"
	"gorm.io/gorm"
)

// Subscriber ...
type Subscriber struct {
	ID                string             `json:"id" gorm:"primaryKey"`
	Mobile            string             `json:"mobile"`
	OfferID           string             `json:"offerID"`
	OfferCode         string             `json:"offerCode"`
	Source            SubscriptionSource `json:"source"`
	Status            SubscriptionStatus `json:"status"`
	StatusDescription string             `json:"statusDescription"`
	CreatedAt         time.Time          `json:"createdAt"`
	UpdatedAt         time.Time          `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt     `json:"deletedAt"`
}

// NewSubscriber ...
func NewSubscriber(mobile, offerID, offerCode string) *Subscriber {
	return &Subscriber{
		ID:                uuid.NewString(),
		Mobile:            mobile,
		OfferID:           offerID,
		OfferCode:         offerCode,
		Source:            SMSSubscription,
		Status:            ScheduledSubscription,
		StatusDescription: "Subscription is scheduled to be sent",
	}
}

// NewWebSubscriber ...
func NewWebSubscriber(mobile, offerID, offerCode string) *Subscriber {
	subscriber := NewSubscriber(mobile, offerID, offerCode)
	subscriber.Source = WebSubscription
	return subscriber
}
