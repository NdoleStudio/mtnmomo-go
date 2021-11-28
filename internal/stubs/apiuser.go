package stubs

// APIUserCreateAPIKey is a dummy json response for the {baseURL}/apiuser/{APIUser}/apikey API
func APIUserCreateAPIKey() []byte {
	return []byte(`
		{
        	"apiKey": "f1db798c98df4bcf83b538175893bbf0"
        }
`)
}

// APIUserGet is a dummy json response for the /v1_0/apiuser/{X-Reference-Id} API
func APIUserGet() []byte {
	return []byte(`
	{
		"providerCallbackHost": "string",
		"targetEnvironment": "sandbox"
	}
`)
}
