package stringsvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	TitleCaseEndpoint        endpoint.Endpoint
	RemoveWhitespaceEndpoint endpoint.Endpoint
	CountEndpoint            endpoint.Endpoint
}

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

func MakeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(ctx, req.S)
		return countResponse{v}, nil
	}
}

func (e Endpoints) TitleCase(ctx context.Context, s string) (string, error) {
	req := titleCaseRequest{S: s}
	res, err := e.TitleCaseEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	return res.(titleCaseResponse).V, nil
}

func (e Endpoints) RemoveWhitespace(ctx context.Context, s string) (string, error) {
	req := removeWhitespaceRequest{S: s}
	res, err := e.RemoveWhitespaceEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	return res.(removeWhitespaceResponse).V, nil
}

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