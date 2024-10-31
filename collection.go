package mtnmomo

import "strings"

type AccountHolderIDType string

const (
	// AccountHolderIDTypeMSISDN is the account holder ID type for a mobile number
	AccountHolderIDTypeMSISDN AccountHolderIDType = "msisdn"

	// AccountHolderIDTypeEMAIL is the account holder ID type for an email address
	AccountHolderIDTypeEMAIL AccountHolderIDType = "email"
)

// AccountHolderStatus is the status of an account holder
type AccountHolderStatus struct {
	IsActive bool `json:"result"`
}

// BasicUserInfo contains personal information of an account holder
type BasicUserInfo struct {
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Sub        string `json:"sub"`
}

// FullName returns the full name of the account holder
func (info *BasicUserInfo) FullName() string {
	return strings.TrimSpace(info.FamilyName + " " + info.GivenName)
}

// RequestToPayParams is the set of parameters used when creating a payment request
type RequestToPayParams struct {
	Amount       string         `json:"amount"`
	Currency     string         `json:"currency"`
	ExternalID   string         `json:"externalId"`
	Payer        *AccountHolder `json:"payer"`
	PayerMessage string         `json:"payerMessage"`
	PayeeNote    string         `json:"payeeNote"`
}

// CollectionTransactionStatus is the status of a request to pay request.
type CollectionTransactionStatus struct {
	Amount                 string         `json:"amount"`
	Currency               string         `json:"currency"`
	ExternalID             string         `json:"externalId"`
	ReferenceID            string         `json:"referenceId"`
	Payer                  *AccountHolder `json:"payer"`
	Status                 string         `json:"status"`
	FinancialTransactionID *string        `json:"financialTransactionId,omitempty"`
	Reason                 *string        `json:"reason,omitempty"`
}

// IsPending checks if a transaction is in pending status
func (status *CollectionTransactionStatus) IsPending() bool {
	return status.Status == "PENDING"
}

// IsFailed checks if a transaction is in failed status
func (status *CollectionTransactionStatus) IsFailed() bool {
	return status.Status == "FAILED"
}

// IsCancelled checks if a transaction is cancelled
func (status *CollectionTransactionStatus) IsCancelled() bool {
	return status.Status == "CANCELLED"
}

// IsSuccessful checks if a transaction is successful
func (status *CollectionTransactionStatus) IsSuccessful() bool {
	return status.Status == "SUCCESSFUL"
}

// AccountHolder identifies an account holder in the wallet platform.
type AccountHolder struct {
	PartyIDType string `json:"partyIdType"`
	PartyID     string `json:"partyId"`
}

// AccountBalance is available balance of the account
type AccountBalance struct {
	AvailableBalance string `json:"availableBalance"`
	Currency         string `json:"currency"`
}
