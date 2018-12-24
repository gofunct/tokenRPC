package server

import (
	"context"

	"github.com/gofunct/hack/pkg/hackserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "testapp/api"
)

// CorgeServiceServer is a composite interface of api_pb.CorgeServiceServer and hackserver.Server.
type CorgeServiceServer interface {
	api_pb.CorgeServiceServer
	hackserver.Server
}

// NewCorgeServiceServer creates a new CorgeServiceServer instance.
func NewCorgeServiceServer() CorgeServiceServer {
	return &corgeServiceServerImpl{}
}

type corgeServiceServerImpl struct {
}

