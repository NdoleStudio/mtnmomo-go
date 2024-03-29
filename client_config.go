package mtnmomo

import "net/http"

type clientConfig struct {
	httpClient        *http.Client
	baseURL           string
	subscriptionKey   string
	targetEnvironment string
	apiKey            string
	apiUser           string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		baseURL:    "https://sandbox.momodeveloper.mtn.com/",
	}
}
