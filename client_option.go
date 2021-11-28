package mtnmomo

import (
	"net/http"
	"strings"
)

// Option is options for constructing a client
type Option interface {
	apply(config *clientConfig)
}

type clientOptionFunc func(config *clientConfig)

func (fn clientOptionFunc) apply(config *clientConfig) {
	fn(config)
}

// WithHTTPClient sets the underlying HTTP client used for API requests.
// By default, http.DefaultClient is used.
func WithHTTPClient(httpClient *http.Client) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if httpClient != nil {
			config.httpClient = httpClient
		}
	})
}

// WithBaseURL set's the base url for the flutterwave API
func WithBaseURL(baseURL string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if baseURL != "" {
			config.baseURL = strings.TrimRight(baseURL, "/")
		}
	})
}

// WithSubscriptionKey sets the delay in milliseconds before a response is gotten.
func WithSubscriptionKey(subscriptionKey string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.subscriptionKey = subscriptionKey
	})
}

// WithAPIUser sets the delay in milliseconds before a response is gotten.
func WithAPIUser(apiUser string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.apiUser = apiUser
	})
}

// WithTargetEnvironment sets the identifier of the EWP system where the transaction shall be processed.
func WithTargetEnvironment(targetEnvironment string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.targetEnvironment = targetEnvironment
	})
}

// WithAPIKey sets the delay in milliseconds before a response is gotten.
func WithAPIKey(apiKey string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.apiKey = apiKey
	})
}
