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

	service.client.addReferenceID(request, userID)

	response, err := service.client.do(request)
	return userID, response, err
}

// CreateAPIKey Used to create an API key for an API user in the sandbox target environment.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/sandbox-provisioning-api/operations/post-v1_0-apiuser-apikey
func (service *apiUserService) CreateAPIKey(ctx context.Context, userID string) (string, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/v1_0/apiuser/"+userID+"/apikey", nil)
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

// Get Used to get API user information.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/sandbox-provisioning-api/operations/get-v1_0-apiuser?
func (service *apiUserService) Get(ctx context.Context, userID string) (*APIUser, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/v1_0/apiuser/"+userID, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	apiUser := new(APIUser)
	if err = json.Unmarshal(*response.Body, apiUser); err != nil {
		return nil, response, err
	}

	apiUser.UserID = userID
	return apiUser, response, err
}
