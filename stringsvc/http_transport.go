package stringsvc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MakeHTTPHandler(endpoints Endpoints, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
	}
	m := http.NewServeMux()
	m.Handle("/tc",
		httptransport.NewServer(
			endpoints.TitleCaseEndpoint,
			DecodeTitleCaseRequest,
			EncodeResponse,
			options...,
		))
	m.Handle("/rw",
		httptransport.NewServer(
			endpoints.RemoveWhitespaceEndpoint,
			DecodeRemoveWhitespaceRequest,
			EncodeResponse,
			options...,
		))
	m.Handle("/c",
		httptransport.NewServer(
			endpoints.CountEndpoint,
			DecodeCountRequest,
			EncodeResponse,
			options...,
		))
	m.Handle("/metrics", promhttp.Handler())
	return m
}

func DecodeTitleCaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request titleCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRemoveWhitespaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request removeWhitespaceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type titleCaseRequest struct {
	S string `json:"s"`
}

type titleCaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type removeWhitespaceRequest struct {
	S string `json:"s"`
}

type removeWhitespaceResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}
