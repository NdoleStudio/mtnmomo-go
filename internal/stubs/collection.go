package stubs

// CollectionToken is a dummy json response for the `/collection/token/` endpoint
func CollectionToken() []byte {
	return []byte(`
		{
			"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			"token_type": "access_token",
			"expires_in": 3600
        }
`)
}
