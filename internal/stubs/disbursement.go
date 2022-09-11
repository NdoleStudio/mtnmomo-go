package stubs

// DisbursementToken is a dummy json response for the `/disbursement/token/` endpoint
func DisbursementToken() []byte {
	return []byte(`
		{
			"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			"token_type": "access_token",
			"expires_in": 3600
        }
`)
}

// DisbursementTransferStatus is a dummy json response for the `/disbursement/v1_0/transfer/{referenceId}` endpoint
func DisbursementTransferStatus() []byte {
	return []byte(`
		{
			"amount": "100",
			"currency": "UGX",
			"financialTransactionId": "23503452",
			"externalId": "947354",
			"payee": {
				"partyIdType": "MSISDN",
				"partyId": "4656473839"
			},
			"status": "SUCCESSFUL"
		}
`)
}

// DisbursementAccountBalance is a dummy json response for the `/disbursement/v1_0/account/balance` endpoint
func DisbursementAccountBalance() []byte {
	return []byte(`
		{
			"availableBalance": "1000",
			"currency": "EUR"
		}
`)
}
