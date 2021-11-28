package mtnmomo

import (
	"context"
	"net/http"
)

// apiUserService is the API client for the `/` endpoint
type apiUserService service

// Create Used to create an API user in the sandbox target environment.
//
// API Docs: https://momodeveloper.mtn.com/docs/services/sandbox-provisioning-api/operations/post-v1_0-apiuser
func (service *apiUserService) Create(ctx context.Context, userID string, providerCallbackHost string) (*Response, error) {
	payload := map[string]string{
		"providerCallbackHost": providerCallbackHost,
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/v1_0/apiuser", payload)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Reference-Id", userID)

	return service.client.do(request)
}
