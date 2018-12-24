package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gofunct/hack/pkg/hackserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/gofunct/hack/pkg/hackserver/testing/api"
)

// NewLibraryServiceServer creates a new LibraryServiceServer instance.
func NewLibraryServiceServer() interface {
	api_pb.LibraryServiceServer
	hackserver.Server
} {
	return &libraryServiceServerImpl{}
}

type libraryServiceServerImpl struct {
}

func (s *libraryServiceServerImpl) ListBooks(ctx context.Context, req *api_pb.ListBooksRequest) (*api_pb.ListBooksResponse, error) {
	return &api_pb.ListBooksResponse{
		Books: []*api_pb.Book{
			{BookId: "The Go Programming Language"},
			{BookId: "Programming Ruby"},
		},
	}, nil
}

func (s *libraryServiceServerImpl) GetBook(ctx context.Context, req *api_pb.GetBookRequest) (*api_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *libraryServiceServerImpl) CreateBook(ctx context.Context, req *api_pb.CreateBookRequest) (*api_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *libraryServiceServerImpl) UpdateBook(ctx context.Context, req *api_pb.UpdateBookRequest) (*api_pb.Book, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}

func (s *libraryServiceServerImpl) DeleteBook(ctx context.Context, req *api_pb.DeleteBookRequest) (*empty.Empty, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
