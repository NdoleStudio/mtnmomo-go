package mtnmomo

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testSubscriptionKey   = "subscriptionKey"
	testAPIKey            = "apiKey"
	testTargetEnvironment = "targetEnvironment"
	testAPIUser           = "apiUser"
)

func TestWithHTTPClient(t *testing.T) {
	t.Run("httpClient is not set when the httpClient is nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithHTTPClient(nil).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
	})

	t.Run("httpClient is set when the httpClient is not nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		newClient := &http.Client{Timeout: 300}

		// Act
		WithHTTPClient(newClient).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
		assert.Equal(t, newClient.Timeout, config.httpClient.Timeout)
	})
}

func TestWithBaseURL(t *testing.T) {
	t.Run("baseURL is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, config.baseURL, config.baseURL)
	})

	t.Run("tailing / is trimmed from baseURL", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com/"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, "https://example.com", config.baseURL)
	})
}

func TestWithSubscriptionKey(t *testing.T) {
	t.Run("subscriptionKey is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithSubscriptionKey(testSubscriptionKey).apply(config)

		// Assert
		assert.Equal(t, testSubscriptionKey, config.subscriptionKey)
	})
}

func TestWithAPIKey(t *testing.T) {
	t.Run("apiKey is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithAPIKey(testAPIKey).apply(config)

		// Assert
		assert.Equal(t, testAPIKey, config.apiKey)
	})
}

func TestWithAPIUser(t *testing.T) {
	t.Run("apiUser is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithAPIUser(testAPIUser).apply(config)

		// Assert
		assert.Equal(t, testAPIUser, config.apiUser)
	})
}

func TestWithTargetEnvironment(t *testing.T) {
	t.Run("targetEnvironment is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithTargetEnvironment(testTargetEnvironment).apply(config)

		// Assert
		assert.Equal(t, testTargetEnvironment, config.targetEnvironment)
	})
}
