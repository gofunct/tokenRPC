package server

import (
	"context"

	"github.com/gofunct/hack/pkg/hackserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	out_pb "testapp/api/out"
)

// QuuxServiceServer is a composite interface of out_pb.QuuxServiceServer and hackserver.Server.
type QuuxServiceServer interface {
	out_pb.QuuxServiceServer
	hackserver.Server
}

// NewQuuxServiceServer creates a new QuuxServiceServer instance.
func NewQuuxServiceServer() QuuxServiceServer {
	return &quuxServiceServerImpl{}
}

type quuxServiceServerImpl struct {
}

