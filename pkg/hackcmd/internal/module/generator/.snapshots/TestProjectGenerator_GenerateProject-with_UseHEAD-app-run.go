package app

import (
	"github.com/gofunct/hack/pkg/hackserver"
)

// Run starts the hackserver.
func Run() error {
	s := hackserver.New(
		hackserver.WithDefaultLogger(),
		hackserver.WithServers(
		// TODO
		),
	)
	return s.Serve()
}

