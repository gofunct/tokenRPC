package server

import (
	"context"

	"github.com/gofunct/hack/pkg/hackserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "testapp/api"
)

// FooServiceServer is a composite interface of api_pb.FooServiceServer and hackserver.Server.
type FooServiceServer interface {
	api_pb.FooServiceServer
	hackserver.Server
}

// NewFooServiceServer creates a new FooServiceServer instance.
func NewFooServiceServer() FooServiceServer {
	return &fooServiceServerImpl{}
}

type fooServiceServerImpl struct {
}

