package mtnmomo

import "net/http"

type clientConfig struct {
	httpClient          *http.Client
	baseURL             string
	subscriptionKey     string
	targetEnvironment   string
	collectionAccount   *apiAccount
	disbursementAccount *apiAccount
}

type apiAccount struct {
	apiUser string
	apiKey  string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		baseURL:    "https://sandbox.momodeveloper.mtn.com/",
	}
}
