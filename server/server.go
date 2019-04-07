package main

import (
	"context"
	"helloworld"
	"net/http"
)

type server struct{}

func (server) Hello(ctx context.Context, req *helloworld.HelloReq) (*helloworld.HelloResp, error) {
	return &helloworld.HelloResp{Text: "Hi there"}, nil
}

func main() {
	http.ListenAndServe(":9090", helloworld.NewHelloWorldServer(server{}, nil))
}
