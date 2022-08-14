package mtnmomo

// RequestToPayParams is the set of parameters used when creating a payment request
type RequestToPayParams struct {
	Amount       string         `json:"amount"`
	Currency     string         `json:"currency"`
	ExternalID   string         `json:"externalId"`
	Payer        *AccountHolder `json:"payer"`
	PayerMessage string         `json:"payerMessage"`
	PayeeNote    string         `json:"payeeNote"`
}

// RequestToPayStatus is the status of a request to pay request.
type RequestToPayStatus struct {
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
func (status *RequestToPayStatus) IsPending() bool {
	return status.Status == "PENDING"
}

// IsFailed checks if a transaction is in failed status
func (status *RequestToPayStatus) IsFailed() bool {
	return status.Status == "FAILED"
}

// IsCancelled checks if a transaction is cancelled
func (status *RequestToPayStatus) IsCancelled() bool {
	return status.Status == "CANCELLED"
}

// IsSuccessful checks if a transaction is successful
func (status *RequestToPayStatus) IsSuccessful() bool {
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
