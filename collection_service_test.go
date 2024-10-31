package mtnmomo

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/NdoleStudio/mtnmomo-go/internal/helpers"
	"github.com/NdoleStudio/mtnmomo-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestCollectionService_Token(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.CollectionToken()}, &requests)
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

	assert.GreaterOrEqual(t, 1, len(requests))
	request := requests[0]
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

func TestCollectionService_RequestToPay(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.CollectionToken(), nil}, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithSubscriptionKey(testSubscriptionKey),
		WithAPIUser(testAPIUser),
		WithAPIKey(testAPIKey),
	)

	// Act
	response, err := client.Collection.RequestToPay(
		context.Background(),
		uuid.NewString(),
		&RequestToPayParams{
			Amount:     "10",
			Currency:   "EUR",
			ExternalID: uuid.NewString(),
			Payer: &AccountHolder{
				PartyIDType: "MSISDN",
				PartyID:     "46733123453",
			},
			PayerMessage: "Test Payer Message",
			PayeeNote:    "Test Payee Note",
		},
		nil,
	)

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 1)
	request := requests[len(requests)-1]
	assert.Equal(t, "/collection/v1_0/requesttopay", request.URL.Path)
	assert.True(t, strings.HasPrefix(request.Header.Get("Authorization"), "Bearer"))
	assert.Equal(t, testSubscriptionKey, request.Header.Get(headerKeySubscriptionKey))
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestCollectionService_GetRequestToPayStatus(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.CollectionToken(), stubs.CollectionRequestToPayStatus()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithSubscriptionKey(testSubscriptionKey),
		WithAPIUser(testAPIUser),
		WithAPIKey(testAPIKey),
	)
	referenceID := uuid.NewString()

	// Act
	status, response, err := client.Collection.GetRequestToPayStatus(context.Background(), referenceID)

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 1)
	request := requests[len(requests)-1]
	assert.Equal(t, "/collection/v1_0/requesttopay/"+referenceID, request.URL.Path)
	assert.True(t, strings.HasPrefix(request.Header.Get("Authorization"), "Bearer"))
	assert.Equal(t, testSubscriptionKey, request.Header.Get(headerKeySubscriptionKey))
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	financialTransactionID := "23503452"
	assert.Equal(t, &CollectionTransactionStatus{
		Amount:                 "100",
		Currency:               "UGX",
		FinancialTransactionID: &financialTransactionID,
		ExternalID:             "947354",
		ReferenceID:            referenceID,
		Payer: &AccountHolder{
			PartyIDType: "MSISDN",
			PartyID:     "4656473839",
		},
		Status: "SUCCESSFUL",
	}, status)

	// Teardown
	server.Close()
}

func TestCollectionService_GetAccountBalance(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.CollectionToken(), stubs.CollectionAccountBalance()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithSubscriptionKey(testSubscriptionKey),
		WithAPIUser(testAPIUser),
		WithAPIKey(testAPIKey),
	)

	// Act
	status, response, err := client.Collection.GetAccountBalance(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 1)
	request := requests[len(requests)-1]
	assert.Equal(t, "/collection/v1_0/account/balance", request.URL.Path)
	assert.True(t, strings.HasPrefix(request.Header.Get("Authorization"), "Bearer"))
	assert.Equal(t, testSubscriptionKey, request.Header.Get(headerKeySubscriptionKey))
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	assert.Equal(t, &AccountBalance{AvailableBalance: "1000", Currency: "EUR"}, status)

	// Teardown
	server.Close()
}

func TestCollectionService_ValidateAccountHolderStatus(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.CollectionToken(), stubs.CollectionValidateAccountHolderStatus()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithSubscriptionKey(testSubscriptionKey),
		WithAPIUser(testAPIUser),
		WithAPIKey(testAPIKey),
	)

	// Act
	status, response, err := client.Collection.ValidateAccountHolderStatus(context.Background(), AccountHolderIDTypeMSISDN, "23777777777")

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 1)
	request := requests[len(requests)-1]
	assert.Equal(t, "/collection/v1_0/accountholder/msisdn/23777777777/active", request.URL.Path)
	assert.True(t, strings.HasPrefix(request.Header.Get("Authorization"), "Bearer"))
	assert.Equal(t, testSubscriptionKey, request.Header.Get(headerKeySubscriptionKey))
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	assert.True(t, status.IsActive)

	// Teardown
	server.Close()
}

func TestCollectionService_ValidateAccountHolderStatusWithInternalServerError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, _, err := client.Collection.ValidateAccountHolderStatus(context.Background(), AccountHolderIDTypeMSISDN, "23777777777")

	// Assert
	assert.NotNil(t, err)

	// Teardown
	server.Close()
}

func TestCollectionService_GetBasicUserinfo(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.CollectionToken(), stubs.CollectionGetBasicUserinfo()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithSubscriptionKey(testSubscriptionKey),
		WithAPIUser(testAPIUser),
		WithAPIKey(testAPIKey),
	)

	// Act
	userInfo, response, err := client.Collection.GetBasicUserinfo(context.Background(), AccountHolderIDTypeMSISDN, "23777777777")

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 1)
	request := requests[len(requests)-1]
	assert.Equal(t, "/collection/v1_0/accountholder/msisdn/23777777777/basicuserinfo", request.URL.Path)
	assert.True(t, strings.HasPrefix(request.Header.Get("Authorization"), "Bearer"))
	assert.Equal(t, testSubscriptionKey, request.Header.Get(headerKeySubscriptionKey))
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	assert.Equal(t, "JOHN DOE", userInfo.FullName())

	// Teardown
	server.Close()
}

func TestCollectionService_GetBasicUserinfoWithInternalServerError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, nil)
	client := New(WithBaseURL(server.URL))

	// Act
	_, _, err := client.Collection.GetBasicUserinfo(context.Background(), AccountHolderIDTypeMSISDN, "23777777777")

	// Assert
	assert.NotNil(t, err)

	// Teardown
	server.Close()
}
