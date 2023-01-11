package models

// SenderNameType ...
type SenderNameType string

// SenderNameTypes ...
const (
	TransactionalSenderName SenderNameType = "Transactional"
	PromotionalSenderName   SenderNameType = "Promotional"
)

// OfferType ...
type OfferType string

// OfferTypes ...
const (
	OnDemandOffer     OfferType = "On-demand"
	SubscriptionOffer OfferType = "Subscription" // recurring
)

// CampaignType ...
type CampaignType string

// CampaignTypes ...
const (
	PremiumCampaign CampaignType = "Premium"
	BulkCampaign    CampaignType = "Bulk"
	DirectCampaign  CampaignType = "Direct"
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

// CampaignMessageType ...
type CampaignMessageType string

// CampaignMessageTypes ...
const (
	PersonalizedCampaignMessage CampaignMessageType = "Personalized"
	BroadcastCampaignMessage    CampaignMessageType = "Broadcast"
)

// SubscriptionSource ...
type SubscriptionSource string

// SubscriptionSources ...
const (
	SMSSubscriptionSource SubscriptionSource = "SMS"
	WebSubscriptionSource SubscriptionSource = "Web"
)

// OutboxSource ...
type OutboxSource string

// SmsSources ...
const (
	APIOutboxSource OutboxSource = "API"
	WebOutboxSource OutboxSource = "Web"
)

// OutboxStatus ...
type OutboxStatus string

// OutboxStatuses ...
const (
	PendingOutboxStatus  OutboxStatus = "Pending"
	SentOutboxStatus     OutboxStatus = "Sent"
	FailedOutboxStatus   OutboxStatus = "Failed"
	CanceledOutboxStatus OutboxStatus = "Canceled"
)

// SubscriberStatus ...
type SubscriberStatus string

// SubscriberStatus ...
const (
	PendingSubscriber    SubscriberStatus = "Pending"
	ProcessingSubscriber SubscriberStatus = "Processing"
	FailedSubscriber     SubscriberStatus = "Failed"
	DeclinedSubscriber   SubscriberStatus = "Declined"
	ActiveSubscriber     SubscriberStatus = "Active"
	Unsubscribed         SubscriberStatus = "Unsubscribed"
	Blacklisted          SubscriberStatus = "Blacklisted"
)
