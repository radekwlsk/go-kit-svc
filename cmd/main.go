package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	grpcclient "github.com/afrometal/go-kit-svc/client/grpc"
	httpclient "github.com/afrometal/go-kit-svc/client/http"
	"github.com/afrometal/go-kit-svc/stringsvc"
	"google.golang.org/grpc"
)

func main() {
	var (
		httpAddr = flag.String("http-addr", "",
			"HTTP address")
		grpcAddr = flag.String("grpc-addr", "",
			"gRPC address")
	)
	flag.Parse()

	var stringService stringsvc.StringService

	if *httpAddr != "" {
		stringService = httpclient.New(*httpAddr)
	} else if *grpcAddr != "" {
		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(),
			grpc.WithTimeout(time.Second))

		if err != nil {
			log.Fatalln("gRPC dial error:", err)
		}
		defer conn.Close()

		stringService = grpcclient.New(conn)
	}

	args := flag.Args()
	var cmd string

	for len(args) > 0 {
		cmd, args = pop(args)
		switch cmd {
		case "tc":
			var s string
			s, args = pop(args)
			titleCase(context.Background(), stringService, s)
		case "rw":
			var s string
			s, args = pop(args)
			removeWhitespace(context.Background(), stringService, s)
		case "c":
			var s string
			s, args = pop(args)
			count(context.Background(), stringService, s)
		default:
			log.Fatalln("unknown command", cmd)
		}
	}
}
func count(ctx context.Context, service stringsvc.StringService, s string) {
	fmt.Println(service.Count(ctx, s))
}
func removeWhitespace(ctx context.Context, service stringsvc.StringService, s string) {
	output, err := service.RemoveWhitespace(ctx, s)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(output)
}
func titleCase(ctx context.Context, service stringsvc.StringService, s string) {
	output, err := service.TitleCase(ctx, s)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(output)
}

// parse command line argument one by one
func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}
