package models

// CampaignChannel ...
type CampaignChannel string

// Channels ...
const (
	BulkSMS    CampaignChannel = "BulkSMS"
	PremiumSMS CampaignChannel = "PremiumSMS"
)

// CampaignType is the type of campaign
type CampaignType string

// CampaignTypes ...
const (
	Transactional CampaignType = "Transactional"
	Promotional   CampaignType = "Promotional"
)

// CampaignScheme  is the scheme of campaign
type CampaignScheme string

// Schemes ...
const (
	Personalized CampaignScheme = "Personalized"
	Broadcast    CampaignScheme = "Broadcast"
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
