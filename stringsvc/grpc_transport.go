package stringsvc

import (
	"context"

	"github.com/afrometal/go-kit-svc/stringsvc/proto"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	oldcontext "golang.org/x/net/context"
)

func MakeGRPCServer(endpoints Endpoints, logger log.Logger) proto.StringServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcServer{
		titleCase: grpctransport.NewServer(
			endpoints.TitleCaseEndpoint,
			DecodeGRPCTitleCaseRequest,
			EncodeGRPCTitleCaseResponse,
			options...,
		),
		removeWhitespace: grpctransport.NewServer(
			endpoints.RemoveWhitespaceEndpoint,
			DecodeGRPCRemoveWhitespaceRequest,
			EncodeGRPCRemoveWhitespaceResponse,
			options...,
		),
		count: grpctransport.NewServer(
			endpoints.CountEndpoint,
			DecodeGRPCCountRequest,
			EncodeGRPCCountResponse,
			options...,
		),
	}
}

type grpcServer struct {
	titleCase        grpctransport.Handler
	removeWhitespace grpctransport.Handler
	count            grpctransport.Handler
}

func (s *grpcServer) TitleCase(ctx oldcontext.Context, req *proto.TitleCaseRequest) (*proto.TitleCaseResponse, error) {
	_, rep, err := s.titleCase.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.TitleCaseResponse), nil
}

func (s *grpcServer) RemoveWhitespace(ctx oldcontext.Context, req *proto.RemoveWhitespaceRequest) (*proto.RemoveWhitespaceResponse, error) {
	_, rep, err := s.removeWhitespace.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.RemoveWhitespaceResponse), nil
}

func (s *grpcServer) Count(ctx oldcontext.Context, req *proto.CountRequest) (*proto.CountResponse, error) {
	_, rep, err := s.count.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.CountResponse), nil
}

func DecodeGRPCTitleCaseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.TitleCaseRequest)
	return titleCaseRequest{S: req.S}, nil
}

func DecodeGRPCRemoveWhitespaceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.RemoveWhitespaceRequest)
	return removeWhitespaceRequest{S: req.S}, nil
}

func DecodeGRPCCountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.CountRequest)
	return countRequest{S: req.S}, nil
}

func DecodeGRPCTitleCaseResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.TitleCaseResponse)
	return titleCaseResponse{V: res.V, Err: res.Err}, nil
}

func DecodeGRPCRemoveWhitespaceResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.RemoveWhitespaceResponse)
	return removeWhitespaceResponse{V: res.V, Err: res.Err}, nil
}

func DecodeGRPCCountResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.CountResponse)
	return countResponse{V: int(res.V)}, nil
}

func EncodeGRPCTitleCaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(titleCaseResponse)
	return &proto.TitleCaseResponse{V: resp.V, Err: resp.Err}, nil
}

func EncodeGRPCRemoveWhitespaceResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(removeWhitespaceResponse)
	return &proto.RemoveWhitespaceResponse{V: resp.V, Err: resp.Err}, nil
}

func EncodeGRPCCountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(countResponse)
	return &proto.CountResponse{V: int64(resp.V)}, nil
}

func EncodeGRPCTitleCaseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(titleCaseRequest)
	return &proto.TitleCaseRequest{S: req.S}, nil
}

func EncodeGRPCRemoveWhitespaceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(removeWhitespaceRequest)
	return &proto.RemoveWhitespaceRequest{S: req.S}, nil
}

func EncodeGRPCCountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(countRequest)
	return &proto.CountRequest{S: req.S}, nil
}
