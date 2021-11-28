package mtnmomo

// AuthToken is A JWT token which is to authorize against the other API end-points.
type AuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}
