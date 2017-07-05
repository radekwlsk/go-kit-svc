package stringsvc

// Server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"

	"github.com/afrometal/go-kit-svc/stringsvc/proto"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	oldcontext "golang.org/x/net/context"
)

// MakeGRPC returns a set of handlers available as a gRPC StringServer.
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

// DecodeGRPCTitleCaseRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request.
// Useful in a server.
func DecodeGRPCTitleCaseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.TitleCaseRequest)
	return titleCaseRequest{S: req.S}, nil
}

// DecodeGRPCRemoveWhitespaceRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request.
// Useful in a server.
func DecodeGRPCRemoveWhitespaceRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.RemoveWhitespaceRequest)
	return removeWhitespaceRequest{S: req.S}, nil
}

// DecodeGRPCCountRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request.
// Useful in a server.
func DecodeGRPCCountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*proto.CountRequest)
	return countRequest{S: req.S}, nil
}

// DecodeGRPCTitleCaseResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC reply to a user-domain response.
// Useful in a client.
func DecodeGRPCTitleCaseResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.TitleCaseResponse)
	return titleCaseResponse{V: res.V, Err: res.Err}, nil
}

// DecodeGRPCRemoveWhitespaceResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC reply to a user-domain response.
// Useful in a client.
func DecodeGRPCRemoveWhitespaceResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.RemoveWhitespaceResponse)
	return removeWhitespaceResponse{V: res.V, Err: res.Err}, nil
}

// DecodeGRPCCountResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC reply to a user-domain response.
// Useful in a client.
func DecodeGRPCCountResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*proto.CountResponse)
	return countResponse{V: int(res.V)}, nil
}

// EncodeGRPCTitleCaseResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply.
// Useful in a server.
func EncodeGRPCTitleCaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(titleCaseResponse)
	return &proto.TitleCaseResponse{V: resp.V, Err: resp.Err}, nil
}

// EncodeGRPCRemoveWhitespaceResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply.
// Useful in a server.
func EncodeGRPCRemoveWhitespaceResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(removeWhitespaceResponse)
	return &proto.RemoveWhitespaceResponse{V: resp.V, Err: resp.Err}, nil
}

// EncodeGRPCCountResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply.
// Useful in a server.
func EncodeGRPCCountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(countResponse)
	return &proto.CountResponse{V: int64(resp.V)}, nil
}

// EncodeGRPCTitleCaseRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain request to a gRPC request.
// Useful in a client.
func EncodeGRPCTitleCaseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(titleCaseRequest)
	return &proto.TitleCaseRequest{S: req.S}, nil
}

// EncodeGRPCRemoveWhitespaceRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain request to a gRPC request.
// Useful in a client.
func EncodeGRPCRemoveWhitespaceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(removeWhitespaceRequest)
	return &proto.RemoveWhitespaceRequest{S: req.S}, nil
}

// EncodeGRPCCountRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain request to a gRPC request.
// Useful in a client.
func EncodeGRPCCountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(countRequest)
	return &proto.CountRequest{S: req.S}, nil
}
