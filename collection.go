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
	Amount                 string              `json:"amount"`
	Currency               string              `json:"currency"`
	ExternalID             string              `json:"externalId"`
	Payer                  *RequestToPayPayer  `json:"payer"`
	Status                 string              `json:"status"`
	FinancialTransactionID *string             `json:"financialTransactionId,omitempty"`
	Reason                 *RequestToPayReason `json:"reason,omitempty"`
}

// RequestToPayPayer identifies an account holder in the wallet platform.
type RequestToPayPayer struct {
	PartyIDType string `json:"partyIdType"`
	PartyID     string `json:"partyId"`
}

// RequestToPayReason contains the cause in case of failure.
type RequestToPayReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// AccountBalance is available balance of the account
type AccountBalance struct {
	AvailableBalance string `json:"availableBalance"`
	Currency         string `json:"currency"`
}
