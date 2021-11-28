package mtnmomo

import (
	"context"
	"encoding/json"
	"net/http"
)

// apiUserService is the API client for the `/` endpoint
type apiUserService service

// CreateAPIUser Used to create an API user in the sandbox target environment.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/sandbox-provisioning-api/operations/post-v1_0-apiuser
func (service *apiUserService) CreateAPIUser(ctx context.Context, userID string, providerCallbackHost string) (string, *Response, error) {
	payload := map[string]string{
		"providerCallbackHost": providerCallbackHost,
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/v1_0/apiuser", payload)
	if err != nil {
		return userID, nil, err
	}
	request.Header.Set("X-Reference-Id", userID)

	response, err := service.client.do(request)
	return userID, response, err
}

// CreateAPIKey Used to create an API user in the sandbox target environment.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/sandbox-provisioning-api/operations/post-v1_0-apiuser
func (service *apiUserService) CreateAPIKey(ctx context.Context, userID string) (string, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/v1_0/apiuser/"+userID+"/apikey ", nil)
	if err != nil {
		return "", nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return "", response, err
	}

	body := struct {
		APIKey string `json:"apiKey"`
	}{}

	if err = json.Unmarshal(*response.Body, &body); err != nil {
		return "", response, err
	}

	return body.APIKey, response, err
}
