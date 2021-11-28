package mtnmomo

import (
	"context"
	"encoding/json"
	"net/http"
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
