package mtnmomo

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// collectionService is the API client for the `/collection` endpoint
type collectionService service

// Token is used to create an access token which can then be used to authorize and authenticate towards the other end-points of the API.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/collection/operations/token-POST
func (service *collectionService) Token(ctx context.Context) (*AuthToken, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/collection/token/", nil)
	if err != nil {
		return nil, nil, err
	}

	service.client.addBasicAuth(request)

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

// RequestToPay is used to request a payment from a consumer (Payer).
//
// API Docs: https://momodeveloper.mtn.com/docs/services/collection/operations/requesttopay-POST
func (service *collectionService) RequestToPay(
	ctx context.Context,
	referenceID string,
	params *RequestToPayParams,
	callbackURL *string,
) (*Response, error) {
	err := service.refreshToken(ctx)
	if err != nil {
		return nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/collection/v1_0/requesttopay", params)
	if err != nil {
		return nil, err
	}

	if callbackURL != nil && len(*callbackURL) != 0 {
		service.client.addCallbackURL(request, *callbackURL)
	}

	service.client.addTargetEnvironment(request)
	service.client.addReferenceID(request, referenceID)
	service.client.addAccessToken(request)

	response, err := service.client.do(request)
	return response, err
}

// GetRequestToPayStatus is used to get the status of a request to pay.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/collection/operations/requesttopay-referenceId-GET
func (service *collectionService) GetRequestToPayStatus(
	ctx context.Context,
	referenceID string,
) (*CollectionTransactionStatus, *Response, error) {
	err := service.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, "/collection/v1_0/requesttopay/"+referenceID, nil)
	if err != nil {
		return nil, nil, err
	}

	service.client.addTargetEnvironment(request)
	service.client.addAccessToken(request)

	response, err := service.client.do(request)
	if err != nil {
		return nil, nil, err
	}

	status := new(CollectionTransactionStatus)
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	status.ReferenceID = referenceID
	return status, response, err
}

// GetAccountBalance returns the balance of the account.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/collection/operations/get-v1_0-account-balance?
func (service *collectionService) GetAccountBalance(ctx context.Context) (*AccountBalance, *Response, error) {
	err := service.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, "/collection/v1_0/account/balance", nil)
	if err != nil {
		return nil, nil, err
	}

	service.client.addAccessToken(request)
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

func (service *collectionService) tokenIsValid() bool {
	return time.Now().UTC().Unix() < service.client.accessTokenExpiresAt
}

func (service *collectionService) refreshToken(ctx context.Context) error {
	service.client.accessTokenLock.Lock()
	defer service.client.accessTokenLock.Unlock()

	if service.tokenIsValid() {
		return nil
	}

	token, _, err := service.Token(ctx)
	if err != nil {
		return err
	}

	service.client.accessToken = token.AccessToken
	service.client.accessTokenExpiresAt = time.Now().UTC().Unix() + token.ExpiresIn - 100

	return nil
}
