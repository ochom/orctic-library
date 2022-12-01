package models

import (
	"io"
	"strconv"
)

// SenderNameType ...
type SenderNameType string

// SenderNameType ...
const (
	Transactional SenderNameType = "Transactional"
	Promotional   SenderNameType = "Promotional"
)

// String ...
func (s SenderNameType) String() string {
	return string(s)
}

// MarshalGQL ...
func (s SenderNameType) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(s.String())))
}

// UnmarshalGQL ...
func (s *SenderNameType) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*s = SenderNameType(v)
		return nil
	default:
		return nil
	}
}

// OfferType ...
type OfferType string

// OfferType ...
const (
	OnDemand     OfferType = "On-demand"
	Subscription OfferType = "Subscription" // recurring
)

// String ...
func (s OfferType) String() string {
	return string(s)
}

// MarshalGQL ...
func (s OfferType) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(s.String())))
}

// UnmarshalGQL ...
func (s *OfferType) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*s = OfferType(v)
		return nil
	default:
		return nil
	}
}

// CampaignType ...
type CampaignType string

const (
	// Premium ...
	Premium CampaignType = "Premium"

	// Bulk ...
	Bulk CampaignType = "Bulk"

	// Direct ...
	Direct CampaignType = "Direct"
)

// String ...
func (s CampaignType) String() string {
	return string(s)
}

// MarshalGQL ...
func (s CampaignType) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(s.String())))
}

// UnmarshalGQL ...
func (s *CampaignType) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*s = CampaignType(v)
		return nil
	default:
		return nil
	}
}

// CampaignMessageType ...
type CampaignMessageType string

// CampaignMessageType ...
const (
	Personalized CampaignMessageType = "Personalized"
	Broadcast    CampaignMessageType = "Broadcast"
)

// String ...
func (s CampaignMessageType) String() string {
	return string(s)
}

// MarshalGQL ...
func (s CampaignMessageType) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(s.String())))
}

// UnmarshalGQL ...
func (s *CampaignMessageType) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*s = CampaignMessageType(v)
		return nil
	default:
		return nil
	}
}

// SubscriptionSource ...
type SubscriptionSource string

// SubscriptionSource ...
const (
	SMS SubscriptionSource = "SMS"
	Web SubscriptionSource = "Web"
)

// String ...
func (s SubscriptionSource) String() string {
	return string(s)
}

// MarshalGQL ...
func (s SubscriptionSource) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(s.String())))
}

// UnmarshalGQL ...
func (s *SubscriptionSource) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*s = SubscriptionSource(v)
		return nil
	default:
		return nil
	}
}
