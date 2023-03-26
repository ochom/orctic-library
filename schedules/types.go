package schedules

// CampaignType ...
type CampaignType string

// Channels ...
const (
	BulkSMS    CampaignType = "BulkSMS"
	PremiumSMS CampaignType = "PremiumSMS"
)

// CampaignSource ...
type CampaignSource string

// CampaignSources ...
const (
	APICampaign CampaignSource = "API"
	WebCampaign CampaignSource = "Web"
)

// OutboxStatus ...
type OutboxStatus string

// OutboxStatuses ...
const (
	PendingOutbox  OutboxStatus = "Pending"
	SentOutbox     OutboxStatus = "Sent"
	FailedOutbox   OutboxStatus = "Failed"
	CanceledOutbox OutboxStatus = "Canceled"
	DeliveryReport OutboxStatus = "DeliveryReport"
)

// SubscriptionSource ...
type SubscriptionSource string

// SubscriptionSources ...
const (
	SMSSubscription SubscriptionSource = "SMS"
	WebSubscription SubscriptionSource = "Web"
)

// SubscriptionStatus ...
type SubscriptionStatus string

// SubscriptionStatus ...
const (
	ScheduledSubscription    SubscriptionStatus = "Scheduled"
	ProcessingSubscription   SubscriptionStatus = "Processing"
	FailedSubscription       SubscriptionStatus = "Failed"
	DeclinedSubscription     SubscriptionStatus = "Declined"
	ActiveSubscription       SubscriptionStatus = "Active"
	UnsubscribedSubscription SubscriptionStatus = "Unsubscribed"
)
