package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	grpcClient "github.com/afrometal/go-kit-svc/client"
	"github.com/afrometal/go-kit-svc/stringsvc"
	"google.golang.org/grpc"
)

// TODO: html client

func main() {
	var (
		grpcAddr = flag.String("addr", ":8081",
			"gRPC address")
	)
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(),
		grpc.WithTimeout(1*time.Second))

	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()

	stringService := grpcClient.New(conn)
	args := flag.Args()
	var cmd string

	for len(args) > 0 {
		cmd, args = pop(args)
		switch cmd {
		case "tc":
			var s string
			s, args = pop(args)
			titleCase(ctx, stringService, s)
		case "rw":
			var s string
			s, args = pop(args)
			removeWhitespace(ctx, stringService, s)
		case "c":
			var s string
			s, args = pop(args)
			count(ctx, stringService, s)
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
