package models

// Payment ...
type Payment struct {
	ID                string        `json:"id"`
	OrganizationID    string        `json:"organizationID"`
	Source            PaymentSource `json:"source"`
	Mobile            string        `json:"mobile"`
	Amount            string        `json:"amount"`
	Status            PaymentStatus `json:"status"`
	StatusDescription string        `json:"statusDescription"`
	ReferenceID       string        `json:"referenceID"`
	BaseModel
}

// Unit ...
type Unit struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationID"`
	PaymentID      string `json:"paymentID"`
	Allocated      int    `json:"allocated"`
	Used           int    `json:"used"`
	BaseModel
}
