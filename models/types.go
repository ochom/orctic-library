package models

import (
	"io"
	"strconv"
)

// SenderNameType ...
type SenderNameType string

// SenderNameType ...
const (
	Transactional SenderNameType = "transactional"
	Promotional   SenderNameType = "promotional"
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
	OnDemand     OfferType = "on-demand"
	Subscription OfferType = "subscription" // recurring
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
	Premium CampaignType = "premium"

	// Bulk ...
	Bulk CampaignType = "bulk"
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
	Personalized CampaignMessageType = "personalized"
	General      CampaignMessageType = "general"
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
