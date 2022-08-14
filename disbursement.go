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
