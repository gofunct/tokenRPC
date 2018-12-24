package server

import (
	"context"

	"github.com/gofunct/hack/pkg/hackserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "testapp/api"
)

// BookServiceServer is a composite interface of api_pb.BookServiceServer and hackserver.Server.
type BookServiceServer interface {
	api_pb.BookServiceServer
	hackserver.Server
}

// NewBookServiceServer creates a new BookServiceServer instance.
func NewBookServiceServer() BookServiceServer {
	return &bookServiceServerImpl{}
}

type bookServiceServerImpl struct {
}

