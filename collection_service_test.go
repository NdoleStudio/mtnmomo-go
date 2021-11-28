package mtnmomo

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/NdoleStudio/mtnmomo-go/internal/helpers"
	"github.com/NdoleStudio/mtnmomo-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestCollectionService_Token(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.CollectionToken()}, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithSubscriptionKey(testSubscriptionKey),
		WithAPIUser(testAPIUser),
		WithAPIKey(testAPIKey),
	)

	// Act
	authToken, response, err := client.Collection.Token(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, 1, len(requests))
	request := requests[0]
	username, password, _ := request.BasicAuth()
	assert.Equal(t, testAPIUser, username)
	assert.Equal(t, testAPIKey, password)
	assert.Equal(t, testSubscriptionKey, request.Header.Get(headerKeySubscriptionKey))
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, int64(3600), authToken.ExpiresIn)
	assert.Equal(t, "access_token", authToken.TokenType)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", authToken.AccessToken)

	// Teardown
	server.Close()
}

func TestCollectionService_RequestToPay(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.CollectionToken(), nil}, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithSubscriptionKey(testSubscriptionKey),
		WithAPIUser(testAPIUser),
		WithAPIKey(testAPIKey),
	)

	// Act
	response, err := client.Collection.RequestToPay(
		context.Background(),
		uuid.NewString(),
		&RequestToPayParams{
			Amount:     "10",
			Currency:   "EUR",
			ExternalID: uuid.NewString(),
			Payer: &RequestToPayParamsPayer{
				PartyIDType: "MSISDN",
				PartyID:     "46733123453",
			},
			PayerMessage: "Test Payer Message",
			PayeeNote:    "Test Payee Note",
		},
		nil,
	)

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 1)
	request := requests[len(requests)-1]
	assert.Equal(t, "/collection/v1_0/requesttopay", request.URL.Path)
	assert.True(t, strings.HasPrefix(request.Header.Get("Authorization"), "Bearer"))
	assert.Equal(t, testSubscriptionKey, request.Header.Get(headerKeySubscriptionKey))
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
