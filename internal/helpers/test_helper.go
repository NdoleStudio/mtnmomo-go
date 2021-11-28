package helpers

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// MakeTestServer creates an api server for testing
func MakeTestServer(responseCode int, body []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(responseCode)
		_, err := res.Write(body)
		if err != nil {
			panic(err)
		}
	}))
}

// MakeRequestCapturingTestServer creates an api server that captures the request object
func MakeRequestCapturingTestServer(responseCode int, responses [][]byte, requests *[]*http.Request) *httptest.Server {
	index := 0
	return httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		clonedRequest := request.Clone(context.Background())

		// clone body
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		request.Body = ioutil.NopCloser(bytes.NewReader(body))
		clonedRequest.Body = ioutil.NopCloser(bytes.NewReader(body))

		*requests = append(*requests, clonedRequest)

		responseWriter.WriteHeader(responseCode)
		_, err = responseWriter.Write(responses[index])
		if err != nil {
			panic(err)
		}
		index++
	}))
}
