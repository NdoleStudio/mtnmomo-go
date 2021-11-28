package mtnmomo

import (
	"context"
	"net/http"
	"testing"

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
	key := "subscriptionKey"
	client := New(WithBaseURL(server.URL), WithSubscriptionKey(key))
	userID := uuid.NewString()

	// Act
	response, err := client.APIUser.Create(context.Background(), userID, "string")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, key, request.Header.Get(subscriptionKeyHeaderKey))
	assert.Equal(t, userID, request.Header.Get("X-Reference-Id"))
	assert.Equal(t, http.StatusCreated, response.HTTPResponse.StatusCode)
}

func TestApiUserService_CreateBadRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusBadRequest, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	response, err := client.APIUser.Create(context.Background(), "errorID", "string")

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.HTTPResponse.StatusCode)
}
