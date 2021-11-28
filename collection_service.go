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

	request, err := service.client.newRequest(ctx, http.MethodPost, "/collection/requesttopay", params)
	if err != nil {
		return nil, err
	}

	if callbackURL != nil && len(*callbackURL) != 0 {
		service.client.addCallbackURL(request, *callbackURL)
	}

	service.client.addTargetEnvironment(request)
	service.client.addReferenceID(request, referenceID)

	response, err := service.client.do(request)
	return response, err
}

func (service *collectionService) tokenIsValid() bool {
	return time.Now().UTC().Unix() < service.client.collectionTokenExpiresAt
}

func (service *collectionService) refreshToken(ctx context.Context) error {
	service.client.collectionLock.Lock()
	defer service.client.collectionLock.Unlock()

	if service.tokenIsValid() {
		return nil
	}

	token, _, err := service.Token(ctx)
	if err != nil {
		return err
	}

	service.client.collectionToken = token.AccessToken
	service.client.collectionTokenExpiresAt = time.Now().UTC().Unix() + token.ExpiresIn - 100

	return nil
}