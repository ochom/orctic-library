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

// PaymentStatus ...
type PaymentStatus string

// PaymentStatuses ...
const (
	PendingPayment   PaymentStatus = "Pending"
	FailedPayment    PaymentStatus = "Failed"
	CanceledPayment  PaymentStatus = "Canceled"
	CompletedPayment PaymentStatus = "Completed"
)

// PaymentSource ...
type PaymentSource string

// PaymentSources ...
const (
	STKPushPaymentSource PaymentSource = "STK Push"
	MpesaPaybillSource   PaymentSource = "Mpesa Paybill"
	BankChequeSource     PaymentSource = "Bank Cheque"
)
