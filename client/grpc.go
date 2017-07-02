package client

import (
	"github.com/afrometal/go-kit-svc/stringsvc"
	"github.com/afrometal/go-kit-svc/stringsvc/proto"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// Return new stringgrpc service
func New(conn *grpc.ClientConn) stringsvc.StringService {
	var titleCaseEndpoint = grpctransport.NewClient(
		conn, "proto.String", "TitleCase",
		stringsvc.EncodeGRPCTitleCaseRequest,
		stringsvc.DecodeGRPCTitleCaseResponse,
		proto.TitleCaseResponse{},
	).Endpoint()

	var removeWhitespaceEndpoint = grpctransport.NewClient(
		conn, "proto.String", "RemoveWhitespace",
		stringsvc.EncodeGRPCRemoveWhitespaceRequest,
		stringsvc.DecodeGRPCRemoveWhitespaceResponse,
		proto.RemoveWhitespaceResponse{},
	).Endpoint()

	var countEndpoint = grpctransport.NewClient(
		conn, "proto.String", "Count",
		stringsvc.EncodeGRPCCountRequest,
		stringsvc.DecodeGRPCCountResponse,
		proto.CountResponse{},
	).Endpoint()

	return stringsvc.Endpoints{
		TitleCaseEndpoint:        titleCaseEndpoint,
		RemoveWhitespaceEndpoint: removeWhitespaceEndpoint,
		CountEndpoint:            countEndpoint,
	}
}
