package mtnmomo

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// disbursementService is the API client for the `/disbursement` endpoint
type disbursementsService service

// Token is used to create an access token which can then be used to authorize and authenticate towards the other end-points of the API.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/disbursement/operations/CreateAccessToken
func (service *disbursementsService) Token(ctx context.Context) (*AuthToken, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/disbursement/token/", nil)
	if err != nil {
		return nil, nil, err
	}

	service.client.addBasicAuth(service.client.disbursementAccount, request)

	response, err := service.client.do(request)
	if err != nil {
		return nil, nil, err
	}

	authToken := new(AuthToken)
	if err = json.Unmarshal(*response.Body, authToken); err != nil {
		return nil, response, err
	}

	return authToken, response, err
}

// Transfer is used to transfer an amount from the owner account to a payee account.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/disbursement/operations/Transfer?
func (service *disbursementsService) Transfer(
	ctx context.Context,
	referenceID string,
	params *TransferParams,
	callbackURL *string,
) (*Response, error) {
	err := service.refreshToken(ctx)
	if err != nil {
		return nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/disbursement/v1_0/transfer", params)
	if err != nil {
		return nil, err
	}

	if callbackURL != nil && len(*callbackURL) != 0 {
		service.client.addCallbackURL(request, *callbackURL)
	}

	service.client.addTargetEnvironment(request)
	service.client.addReferenceID(request, referenceID)
	service.client.addDisbursementAccessToken(request)

	response, err := service.client.do(request)
	return response, err
}

// GetTransferStatus is used to get the status of a transfer.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/disbursement/operations/GetTransferStatus
func (service *disbursementsService) GetTransferStatus(
	ctx context.Context,
	referenceID string,
) (*DisbursementTransactionStatus, *Response, error) {
	err := service.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, "/disbursement/v1_0/transfer/"+referenceID, nil)
	if err != nil {
		return nil, nil, err
	}

	service.client.addTargetEnvironment(request)
	service.client.addDisbursementAccessToken(request)

	response, err := service.client.do(request)
	if err != nil {
		return nil, nil, err
	}

	status := new(DisbursementTransactionStatus)
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	status.ReferenceID = referenceID
	return status, response, err
}

// GetAccountBalance returns the balance of the account.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/disbursement/operations/GetAccountBalance
func (service *disbursementsService) GetAccountBalance(ctx context.Context) (*AccountBalance, *Response, error) {
	err := service.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, "/disbursement/v1_0/account/balance", nil)
	if err != nil {
		return nil, nil, err
	}

	service.client.addDisbursementAccessToken(request)
	service.client.addTargetEnvironment(request)

	response, err := service.client.do(request)
	if err != nil {
		return nil, nil, err
	}

	balance := new(AccountBalance)
	if err = json.Unmarshal(*response.Body, balance); err != nil {
		return nil, response, err
	}

	return balance, response, err
}

func (service *disbursementsService) tokenIsValid() bool {
	return time.Now().UTC().Unix() < service.client.disbursementAccessTokenExpiresAt
}

func (service *disbursementsService) refreshToken(ctx context.Context) error {
	service.client.disbursementLock.Lock()
	defer service.client.disbursementLock.Unlock()

	if service.tokenIsValid() {
		return nil
	}

	token, _, err := service.Token(ctx)
	if err != nil {
		return err
	}

	service.client.disbursementAccessToken = token.AccessToken
	service.client.disbursementAccessTokenExpiresAt = time.Now().UTC().Unix() + token.ExpiresIn - 100

	return nil
}
