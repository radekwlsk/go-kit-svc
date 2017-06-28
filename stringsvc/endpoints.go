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
		v, err := svc.TitleCase(req.S)
		if err != nil {
			return titleCaseResponse{v, err.Error()}, nil
		}
		return titleCaseResponse{v, ""}, nil
	}
}

func MakeRemoveWhitespaceEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(removeWhitespaceRequest)
		v, err := svc.RemoveWhitespace(req.S)
		if err != nil {
			return removeWhitespaceResponse{v, err.Error()}, nil
		}
		return removeWhitespaceResponse{v, ""}, nil
	}
}

func MakeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}
