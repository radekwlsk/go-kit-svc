package stringsvc

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

func MakeHTTPHandler(endpoints Endpoints, logger log.Logger) http.Handler {
	logger = log.With(logger, "transport", "HTTP")
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

func DecodeHTTPTitleCaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request titleCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeHTTPRemoveWhitespaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request removeWhitespaceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeHTTPCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeHTTPTitleCaseResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp titleCaseResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func DecodeHTTPRemoveWhitespaceResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp removeWhitespaceResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func DecodeHTTPCountResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp countResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func EncodeHTTPRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

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
