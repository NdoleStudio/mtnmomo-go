package mtnmomo

// RequestToPayParams is the set of parameters used when creating a payment request
type RequestToPayParams struct {
	Amount       string                   `json:"amount"`
	Currency     string                   `json:"currency"`
	ExternalID   string                   `json:"externalId"`
	Payer        *RequestToPayParamsPayer `json:"payer"`
	PayerMessage string                   `json:"payerMessage"`
	PayeeNote    string                   `json:"payeeNote"`
}

// RequestToPayParamsPayer identifies an account holder in the wallet platform.
type RequestToPayParamsPayer struct {
	PartyIDType string `json:"partyIdTyp"`
	PartyID     string `json:"partyId"`
}
