# String Service

Microservice training using Go, [go-kit](https://gokit.io/) toolkit and [gRPC](http://www.grpc.io/) RPC framework for message transport.

Based on [go-kit examples](https://gokit.io/examples/):
- [stringsvc](https://gokit.io/examples/stringsvc.html)
- [addsvc](https://github.com/go-kit/kit/tree/master/examples/addsvc)

and [ru rocker's tutorial](http://www.ru-rocker.com/2017/02/24/micro-services-using-go-kit-grpc-endpoint/) on microservices using go-kit and gRPC.

StringService provides only 3 basic operations on strings as project is made to learn microservices and gRPC principles more than Go itself.

## Usage

Run server using `main.go` file:
```bash
$ go run main.go
```

Now HTTP requests can be made:
```bash
$ curl -XPOST -d'{"s":"hello, world!"}' localhost:8080/tc
{"v":"Hello, World!"}
$ curl -XPOST -d'{"s":"hello, world!"}' localhost:8080/rw
{"v":"hello,world!"}
$ curl -XPOST -d'{"s":"hello, world!"}' localhost:8080/c
{"v":13}
```

Or gRPC client can be run:
```bash
$ go run cmd/client.go tc "hello, world!" rw "hello,   world!" c "hello, world!"
Hello, World!
hello,world!
13
```
program arguments are in form `cmd str cmd str cmd str ...` where `cmd` can be:

- `tc` for `StringService.TitleCase(...)`
- `rw` for `StringService.RemoveWhitespace(...)`
- `c` for `StringService.Count(...)`

and `str` can be any string that will be argument of command.