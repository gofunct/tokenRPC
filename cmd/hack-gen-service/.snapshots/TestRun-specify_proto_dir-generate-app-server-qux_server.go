package server

import (
	"context"

	"github.com/gofunct/hack/pkg/hackserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "testapp/api"
)

// QuxServiceServer is a composite interface of api_pb.QuxServiceServer and hackserver.Server.
type QuxServiceServer interface {
	api_pb.QuxServiceServer
	hackserver.Server
}

// NewQuxServiceServer creates a new QuxServiceServer instance.
func NewQuxServiceServer() QuxServiceServer {
	return &quxServiceServerImpl{}
}

type quxServiceServerImpl struct {
}

