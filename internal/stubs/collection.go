package stubs

// CollectionToken is a dummy json response for the `/collection/token/` endpoint
func CollectionToken() []byte {
	return []byte(`
		{
			"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			"token_type": "access_token",
			"expires_in": 3600
        }
`)
}

// CollectionRequestToPayStatus is a dummy json response for the `/collection/v1_0/requesttopay/{referenceId}` endpoint
func CollectionRequestToPayStatus() []byte {
	return []byte(`
		{
			"amount": "100",
			"currency": "UGX",
			"financialTransactionId": "23503452",
			"externalId": "947354",
			"payer": {
				"partyIdType": "MSISDN",
				"partyId": "4656473839"
			},
			"status": "SUCCESSFUL"
		}
`)
}

// CollectionAccountBalance is a dummy json response for the `/collection/v1_0/account/balance` endpoint
func CollectionAccountBalance() []byte {
	return []byte(`
		{
			"availableBalance": "1000",
			"currency": "EUR"
		}
`)
}

// CollectionValidateAccountHolderStatus is a dummy json response for the `/collection/v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/active` endpoint
func CollectionValidateAccountHolderStatus() []byte {
	return []byte(`
		{
			"result": true
		}
`)
}

// CollectionGetBasicUserinfo is a dummy json response for the `/collection/v1_0/accountholder/{accountHolderIdType}/{accountHolderId}/basicuserinfo` endpoint
func CollectionGetBasicUserinfo() []byte {
	return []byte(`
		{
			"family_name": "JOHN",
			"given_name": "DOE",
			"sub": "1111111"
		}
`)
}
