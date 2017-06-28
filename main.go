package main

import (
	"net/http"
	"os"

	"./stringsvc"
	"github.com/go-kit/kit/log"
)

func main() {
	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	logger.Log("msg", "hello")
	defer logger.Log("msg", "goodbye")

	var svc stringsvc.StringService
	svc = stringsvc.New()
	svc = stringsvc.NewLoggingMiddleware(svc, logger)
	svc = stringsvc.NewInstrumentingMiddleware(svc)

	var endpoints stringsvc.Endpoints
	endpoints = stringsvc.Endpoints{
		TitleCaseEndpoint:        stringsvc.MakeTitleCaseEndpoint(svc),
		RemoveWhitespaceEndpoint: stringsvc.MakeRemoveWhitespaceEndpoint(svc),
		CountEndpoint:            stringsvc.MakeCountEndpoint(svc),
	}

	logger = log.With(logger, "transport", "HTTP")
	logger.Log("msg", "HTTP", "addr", ":8080")

	handler := stringsvc.MakeHTTPHandler(endpoints, logger)
	http.ListenAndServe(":8080", handler)
}
