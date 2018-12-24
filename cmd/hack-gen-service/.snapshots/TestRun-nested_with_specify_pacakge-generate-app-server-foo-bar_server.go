package server

import (
	"context"

	"github.com/gofunct/hack/pkg/hackserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	foo_pb "testapp/api/foo"
)

// BarServiceServer is a composite interface of foo_pb.BarServiceServer and hackserver.Server.
type BarServiceServer interface {
	foo_pb.BarServiceServer
	hackserver.Server
}

// NewBarServiceServer creates a new BarServiceServer instance.
func NewBarServiceServer() BarServiceServer {
	return &barServiceServerImpl{}
}

type barServiceServerImpl struct {
}

