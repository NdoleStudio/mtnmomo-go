package mtnmomo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	headerKeySubscriptionKey   = "Ocp-Apim-Subscription-Key"
	headerKeyTargetEnvironment = "X-Target-Environment"
	headerKeyReferenceID       = "X-Reference-Id"
	headerKeyCallbackURL       = "X-Callback-Url"
)

type service struct {
	client *Client
}

// Client is the campay API client.
// Do not instantiate this client with Client{}. Use the New method instead.
type Client struct {
	httpClient          *http.Client
	common              service
	baseURL             string
	subscriptionKey     string
	collectionAccount   *apiAccount
	disbursementAccount *apiAccount
	targetEnvironment   string

	collectionLock                 sync.Mutex
	collectionAccessToken          string
	collectionAccessTokenExpiresAt int64

	disbursementLock                 sync.Mutex
	disbursementAccessToken          string
	disbursementAccessTokenExpiresAt int64

	APIUser      *apiUserService
	Collection   *collectionService
	Disbursement *disbursementsService
}

// New creates and returns a new campay.Client from a slice of campay.ClientOption.
func New(options ...Option) *Client {
	config := defaultClientConfig()

	for _, option := range options {
		option.apply(config)
	}

	client := &Client{
		httpClient:          config.httpClient,
		subscriptionKey:     config.subscriptionKey,
		baseURL:             config.baseURL,
		collectionAccount:   config.collectionAccount,
		disbursementAccount: config.disbursementAccount,
		targetEnvironment:   config.targetEnvironment,
		collectionLock:      sync.Mutex{},
	}

	client.common.client = client
	client.APIUser = (*apiUserService)(&client.common)
	client.Collection = (*collectionService)(&client.common)
	client.Disbursement = (*disbursementsService)(&client.common)

	return client
}

// newRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the BaseURL of the Client.
// URI's should always be specified without a preceding slash.
func (client *Client) newRequest(ctx context.Context, method, uri string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, client.baseURL+uri, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(headerKeySubscriptionKey, client.subscriptionKey)

	return req, nil
}

func (client *Client) addCollectionAccessToken(request *http.Request) {
	request.Header.Add("Authorization", "Bearer "+client.collectionAccessToken)
}

func (client *Client) addDisbursementAccessToken(request *http.Request) {
	request.Header.Add("Authorization", "Bearer "+client.disbursementAccessToken)
}

func (client *Client) addBasicAuth(account *apiAccount, request *http.Request) {
	request.SetBasicAuth(account.apiUser, account.apiKey)
}

func (client *Client) addReferenceID(request *http.Request, reference string) {
	request.Header.Set(headerKeyReferenceID, reference)
}

func (client *Client) addCallbackURL(request *http.Request, url string) {
	request.Header.Set(headerKeyCallbackURL, url)
}

func (client *Client) addTargetEnvironment(request *http.Request) {
	request.Header.Set(headerKeyTargetEnvironment, client.targetEnvironment)
}

// do carries out an HTTP request and returns a Response
func (client *Client) do(req *http.Request) (*Response, error) {
	httpResponse, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = httpResponse.Body.Close() }()

	resp, err := client.newResponse(httpResponse)
	if err != nil {
		return resp, err
	}

	_, err = io.Copy(ioutil.Discard, httpResponse.Body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// newResponse converts an *http.Response to *Response
func (client *Client) newResponse(httpResponse *http.Response) (*Response, error) {
	response := new(Response)
	response.HTTPResponse = httpResponse

	buf, err := ioutil.ReadAll(response.HTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	response.Body = &buf

	return response, response.Error()
}
