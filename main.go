package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"flag"

	"github.com/afrometal/go-kit-svc/stringsvc"
	"github.com/afrometal/go-kit-svc/stringsvc/proto"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

func main() {
	var (
		httpAddr = flag.String("http-addr", ":8080",
			"HTTP address")
		grpcAddr = flag.String("grpc-addr", ":8081",
			"gRPC address")
	)
	flag.Parse()
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
	{
		svc = stringsvc.New()
		svc = stringsvc.NewLoggingMiddleware(svc, logger)
		svc = stringsvc.NewInstrumentingMiddleware(svc)
	}
	
	var endpoints stringsvc.Endpoints
	{
		endpoints = stringsvc.Endpoints{
			TitleCaseEndpoint:        stringsvc.MakeTitleCaseEndpoint(svc),
			RemoveWhitespaceEndpoint: stringsvc.MakeRemoveWhitespaceEndpoint(svc),
			CountEndpoint:            stringsvc.MakeCountEndpoint(svc),
		}
	}

	errc := make(chan error)

	go func() {
		logger := log.With(logger, "transport", "HTTP")
		logger.Log("addr", *httpAddr)

		handler := stringsvc.MakeHTTPHandler(endpoints, logger)
		errc <- http.ListenAndServe(*httpAddr, handler)
	}()

	go func() {
		logger := log.With(logger, "transport", "gRPC")
		logger.Log("addr", *grpcAddr)

		ln, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errc <- err
			return
		}

		srv := stringsvc.MakeGRPCServer(endpoints, logger)
		s := grpc.NewServer()
		proto.RegisterStringServer(s, srv)
		errc <- s.Serve(ln)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Run!
	logger.Log("exit", <-errc)
}
