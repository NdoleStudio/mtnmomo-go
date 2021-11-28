package mtnmomo

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/mtnmomo-go/internal/helpers"
	"github.com/NdoleStudio/mtnmomo-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestCollectionService_Token(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	request := http.Request{}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.CollectionToken(), &request)
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
