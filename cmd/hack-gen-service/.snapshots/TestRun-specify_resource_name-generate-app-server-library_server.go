package server

import (
	"context"

	"github.com/gofunct/hack/pkg/hackserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "testapp/api"
)

// LibraryServiceServer is a composite interface of api_pb.LibraryServiceServer and hackserver.Server.
type LibraryServiceServer interface {
	api_pb.LibraryServiceServer
	hackserver.Server
}

// NewLibraryServiceServer creates a new LibraryServiceServer instance.
func NewLibraryServiceServer() LibraryServiceServer {
	return &libraryServiceServerImpl{}
}

type libraryServiceServerImpl struct {
}

