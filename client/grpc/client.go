// Package grpc provides a gRPC client for the string service.
package grpc

import (
	"github.com/afrometal/go-kit-svc/stringsvc"
	"github.com/afrometal/go-kit-svc/stringsvc/proto"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// New returns StringService based on gRPC client connection.
// Caller have to dial and close the connection.
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
