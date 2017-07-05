// Package http provides a HTTP client for the string service.
package http

import (
	"net/url"
	"strings"

	"github.com/afrometal/go-kit-svc/stringsvc"
	httptransport "github.com/go-kit/kit/transport/http"
)

// New returns StringService based on HTTP server at remote instance.
// Instance is expected to come in "host:port" form.
func New(instance string) stringsvc.StringService {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	var titleCaseEndpoint = httptransport.NewClient(
		"POST",
		copyURL(u, "/tc"),
		stringsvc.EncodeHTTPRequest,
		stringsvc.DecodeHTTPTitleCaseResponse,
	).Endpoint()

	var removeWhitespaceEndpoint = httptransport.NewClient(
		"POST",
		copyURL(u, "/rw"),
		stringsvc.EncodeHTTPRequest,
		stringsvc.DecodeHTTPRemoveWhitespaceResponse,
	).Endpoint()

	var countEndpoint = httptransport.NewClient(
		"POST",
		copyURL(u, "/c"),
		stringsvc.EncodeHTTPRequest,
		stringsvc.DecodeHTTPCountResponse,
	).Endpoint()

	return stringsvc.Endpoints{
		TitleCaseEndpoint:        titleCaseEndpoint,
		RemoveWhitespaceEndpoint: removeWhitespaceEndpoint,
		CountEndpoint:            countEndpoint,
	}
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}
