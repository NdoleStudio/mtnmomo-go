package mtnmomo

// RequestToPayParams is the set of parameters used when creating a payment request
type RequestToPayParams struct {
	Amount       string             `json:"amount"`
	Currency     string             `json:"currency"`
	ExternalID   string             `json:"externalId"`
	Payer        *RequestToPayPayer `json:"payer"`
	PayerMessage string             `json:"payerMessage"`
	PayeeNote    string             `json:"payeeNote"`
}

// RequestToPayStatus is the status of a request to pay request.
type RequestToPayStatus struct {
	Amount                 string             `json:"amount"`
	Currency               string             `json:"currency"`
	ExternalID             string             `json:"externalId"`
	Payer                  *RequestToPayPayer `json:"payer"`
	Status                 string             `json:"status"`
	FinancialTransactionID *string            `json:"financialTransactionId,omitempty"`
	Reason                 *string            `json:"reason,omitempty"`
}

// RequestToPayPayer identifies an account holder in the wallet platform.
type RequestToPayPayer struct {
	PartyIDType string `json:"partyIdType"`
	PartyID     string `json:"partyId"`
}

// AccountBalance is available balance of the account
type AccountBalance struct {
	AvailableBalance string `json:"availableBalance"`
	Currency         string `json:"currency"`
}
