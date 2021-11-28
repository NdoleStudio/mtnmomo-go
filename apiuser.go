package mtnmomo

// APIUser contains information about an API User
type APIUser struct {
	UserID               string `json:"userId"`
	ProviderCallbackHost string `json:"providerCallbackHost"`
	TargetEnvironment    string `json:"targetEnvironment"`
}
