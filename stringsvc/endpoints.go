package stringsvc

// Methods to make individual endpoints from services,
// request and response types to serve those endpoints.

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints collects all endpoints that are required by StringService.
// It's a helper struct to collect all of the endpoints into a single parameter.
type Endpoints struct {
	TitleCaseEndpoint        endpoint.Endpoint
	RemoveWhitespaceEndpoint endpoint.Endpoint
	CountEndpoint            endpoint.Endpoint
}

// MakeTitleCaseEndpoint returns an endpoint that invokes TitleCase on the StringService.
// Useful in a server.
func MakeTitleCaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(titleCaseRequest)
		v, err := svc.TitleCase(ctx, req.S)
		if err != nil {
			return titleCaseResponse{v, err.Error()}, nil
		}
		return titleCaseResponse{v, ""}, nil
	}
}

// MakeRemoveWhitespaceEndpoint returns an endpoint that invokes
// RemoveWhitespace on the StringService.
// Useful in a server.
func MakeRemoveWhitespaceEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(removeWhitespaceRequest)
		v, err := svc.RemoveWhitespace(ctx, req.S)
		if err != nil {
			return removeWhitespaceResponse{v, err.Error()}, nil
		}
		return removeWhitespaceResponse{v, ""}, nil
	}
}

// MakeCountEndpoint returns an endpoint that invokes Count on the StringService.
// Useful in a server.
func MakeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(ctx, req.S)
		return countResponse{v}, nil
	}
}

// TitleCase implements StringService.
// Useful in a client.
func (e Endpoints) TitleCase(ctx context.Context, s string) (string, error) {
	req := titleCaseRequest{S: s}
	res, err := e.TitleCaseEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	return res.(titleCaseResponse).V, nil
}

// RemoveWhitespace implements StringService.
// Useful in a client.
func (e Endpoints) RemoveWhitespace(ctx context.Context, s string) (string, error) {
	req := removeWhitespaceRequest{S: s}
	res, err := e.RemoveWhitespaceEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	return res.(removeWhitespaceResponse).V, nil
}

// Count implements StringService.
// Useful in a client.
func (e Endpoints) Count(ctx context.Context, s string) int {
	req := countRequest{S: s}
	res, err := e.CountEndpoint(ctx, req)
	if err != nil {
		return 0
	}
	return res.(countResponse).V
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
