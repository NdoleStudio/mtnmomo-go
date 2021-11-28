package mtnmomo

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/NdoleStudio/mtnmomo-go/internal/stubs"

	"github.com/NdoleStudio/mtnmomo-go/internal/helpers"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestApiUserService_Create(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	request := http.Request{}
	server := helpers.MakeRequestCapturingTestServer(http.StatusCreated, nil, &request)
	client := New(WithBaseURL(server.URL), WithSubscriptionKey(testSubscriptionKey))
	userID := uuid.NewString()

	// Act
	apiUser, response, err := client.APIUser.CreateAPIUser(context.Background(), userID, "string")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, testSubscriptionKey, request.Header.Get(subscriptionKeyHeaderKey))
	assert.Equal(t, userID, request.Header.Get("X-Reference-Id"))
	assert.Equal(t, http.StatusCreated, response.HTTPResponse.StatusCode)
	assert.Equal(t, userID, apiUser)

	// Teardown
	server.Close()
}

func TestApiUserService_CreateBadRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusBadRequest, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.APIUser.CreateAPIUser(context.Background(), "errorID", "string")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestApiUserService_CreateAPIKey(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	request := http.Request{}
	server := helpers.MakeRequestCapturingTestServer(http.StatusCreated, stubs.APIUserCreateAPIKey(), &request)
	client := New(WithBaseURL(server.URL), WithSubscriptionKey(testSubscriptionKey))
	userID := uuid.NewString()

	// Act
	apiKey, response, err := client.APIUser.CreateAPIKey(context.Background(), userID)

	// Assert
	assert.Nil(t, err)

	assert.True(t, strings.Contains(request.URL.String(), userID))
	assert.Equal(t, testSubscriptionKey, request.Header.Get(subscriptionKeyHeaderKey))
	assert.Equal(t, "f1db798c98df4bcf83b538175893bbf0", apiKey)
	assert.Equal(t, http.StatusCreated, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestApiUserService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	request := http.Request{}
	server := helpers.MakeRequestCapturingTestServer(http.StatusCreated, stubs.APIUserGet(), &request)
	client := New(WithBaseURL(server.URL), WithSubscriptionKey(testSubscriptionKey))
	userID := uuid.NewString()

	// Act
	apiUser, response, err := client.APIUser.Get(context.Background(), userID)

	// Assert
	assert.Nil(t, err)

	assert.True(t, strings.Contains(request.URL.String(), userID))
	assert.Equal(t, testSubscriptionKey, request.Header.Get(subscriptionKeyHeaderKey))
	assert.Equal(t, http.StatusCreated, response.HTTPResponse.StatusCode)
	assert.Equal(t, userID, apiUser.UserID)
	assert.Equal(t, "sandbox", apiUser.TargetEnvironment)
	assert.Equal(t, "string", apiUser.ProviderCallbackHost)

	// Teardown
	server.Close()
}
