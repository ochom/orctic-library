package models

// CampaignType ...
type CampaignType string

// CampaignTypes ...
const (
	PersonalizedCampaign CampaignType = "Personalized"
	BroadcastCampaign    CampaignType = "Broadcast"
)

// CampaignSource ...
type CampaignSource string

// CampaignSources ...
const (
	APICampaign CampaignSource = "API"
	WebCampaign CampaignSource = "Web"
)

// CampaignStatus ...
type CampaignStatus string

// CampaignStatuses ...
const (
	PendingCampaign  CampaignStatus = "Pending"
	SentCampaign     CampaignStatus = "Sent"
	FailedCampaign   CampaignStatus = "Failed"
	CanceledCampaign CampaignStatus = "Canceled"
)

// OutboxType ...
type OutboxType string

// OutboxTypes ...
const (
	TransactionalOutbox OutboxType = "Transactional"
	PromotionalOutbox   OutboxType = "Promotional"
)

// OutboxSource ...
type OutboxSource string

// SmsSources ...
const (
	APIOutbox OutboxSource = "API"
	WebOutbox OutboxSource = "Web"
)

// OutboxStatus ...
type OutboxStatus string

// OutboxStatuses ...
const (
	PendingOutbox  OutboxStatus = "Pending"
	SentOutbox     OutboxStatus = "Sent"
	FailedOutbox   OutboxStatus = "Failed"
	CanceledOutbox OutboxStatus = "Canceled"
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
