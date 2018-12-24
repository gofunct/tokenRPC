package server

import (
	"context"

	"github.com/gofunct/hack/pkg/hackserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	foo_pb "testapp/api/foo"
)

// BarBazServiceServer is a composite interface of foo_pb.BarBazServiceServer and hackserver.Server.
type BarBazServiceServer interface {
	foo_pb.BarBazServiceServer
	hackserver.Server
}

// NewBarBazServiceServer creates a new BarBazServiceServer instance.
func NewBarBazServiceServer() BarBazServiceServer {
	return &barBazServiceServerImpl{}
}

type barBazServiceServerImpl struct {
}

