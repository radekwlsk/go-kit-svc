package stringsvc

// Server-side bindings for the HTTP transport.
// It utilizes the transport/http.Server.

import (
	"context"
	"encoding/json"
	"net/http"

	"bytes"
	"errors"
	"io/ioutil"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MakeHTTPHandler returns a handler that makes a set of endpoints available
// on predefined paths.
func MakeHTTPHandler(endpoints Endpoints, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
	}
	m := http.NewServeMux()
	m.Handle("/tc",
		httptransport.NewServer(
			endpoints.TitleCaseEndpoint,
			DecodeHTTPTitleCaseRequest,
			EncodeHTTPResponse,
			options...,
		))
	m.Handle("/rw",
		httptransport.NewServer(
			endpoints.RemoveWhitespaceEndpoint,
			DecodeHTTPRemoveWhitespaceRequest,
			EncodeHTTPResponse,
			options...,
		))
	m.Handle("/c",
		httptransport.NewServer(
			endpoints.CountEndpoint,
			DecodeHTTPCountRequest,
			EncodeHTTPResponse,
			options...,
		))
	m.Handle("/metrics", promhttp.Handler())
	return m
}

// DecodeHTTPTitleCaseRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
// Useful in a server.
func DecodeHTTPTitleCaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request titleCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeHTTPRemoveWhitespaceRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
// Useful in a server.
func DecodeHTTPRemoveWhitespaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request removeWhitespaceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeHTTPCountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
// Useful in a server.
func DecodeHTTPCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeHTTPTitleCaseResponse is a transport/http.DecodeResponseFunc that decodes a
// JSON-encoded response from the HTTP response body. For non-200 status code response
// an error message decoding attempt is made on response body.
// Useful in a client.
func DecodeHTTPTitleCaseResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp titleCaseResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// DecodeHTTPRemoveWhitespaceResponse is a transport/http.DecodeResponseFunc that decodes a
// JSON-encoded response from the HTTP response body. For non-200 status code response
// an error message decoding attempt is made on response body.
// Useful in a client.
func DecodeHTTPRemoveWhitespaceResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp removeWhitespaceResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// DecodeHTTPCountResponse is a transport/http.DecodeResponseFunc that decodes a
// JSON-encoded response from the HTTP response body. For non-200 status code response
// an error message decoding attempt is made on response body.
// Useful in a client.
func DecodeHTTPCountResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp countResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// EncodeHTTPRequest is a transport/http.EncodeRequestFunc that JSON-encodes any
// request to the request body.
// Useful in a client.
func EncodeHTTPRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// EncodeHTTPResponse is a transport/http.EncodeResponseFunc that encodes the response
// as JSON to the response writer.
// Useful in a server.
func EncodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func errorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}
