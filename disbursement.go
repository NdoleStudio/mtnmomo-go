package mtnmomo

// TransferParams is the set of parameters for transferring money to a payee account
type TransferParams struct {
	Amount       string         `json:"amount"`
	Currency     string         `json:"currency"`
	ExternalID   string         `json:"externalId"`
	Payee        *AccountHolder `json:"payee"`
	PayerMessage string         `json:"payerMessage"`
	PayeeNote    string         `json:"payeeNote"`
}

// DisbursementTransactionStatus is the status of a request to pay request.
type DisbursementTransactionStatus struct {
	Amount                 string         `json:"amount"`
	Currency               string         `json:"currency"`
	ExternalID             string         `json:"externalId"`
	ReferenceID            string         `json:"referenceId"`
	Payee                  *AccountHolder `json:"payee"`
	Status                 string         `json:"status"`
	FinancialTransactionID *string        `json:"financialTransactionId,omitempty"`
	PayerMessage           string         `json:"payerMessage"`
	PayeeNote              string         `json:"payeeNote"`
}

// IsPending checks if a transaction is in pending status
func (status *DisbursementTransactionStatus) IsPending() bool {
	return status.Status == "PENDING"
}

// IsFailed checks if a transaction is in failed status
func (status *DisbursementTransactionStatus) IsFailed() bool {
	return status.Status == "FAILED"
}

// IsCancelled checks if a transaction is cancelled
func (status *DisbursementTransactionStatus) IsCancelled() bool {
	return status.Status == "CANCELLED"
}

// IsSuccessful checks if a transaction is successful
func (status *DisbursementTransactionStatus) IsSuccessful() bool {
	return status.Status == "SUCCESSFUL"
}
