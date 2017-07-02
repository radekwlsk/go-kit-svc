package html

import (
	"log"
	"net/url"
	"strings"

	"github.com/afrometal/go-kit-svc/stringsvc"
	httptransport "github.com/go-kit/kit/transport/http"
)

func New(instance string, logger log.Logger) stringsvc.StringService {
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
