package svcgen

import (
	"github.com/gofunct/hack/pkg/gencmd"
	"github.com/gofunct/hack/pkg/protoc"
	"github.com/gofunct/hack/pkg/svcgen/params"
)

type CreateAppFunc func(*gencmd.Command) (*App, error)

type App struct {
	ProtocWrapper protoc.Wrapper
	ParamsBuilder params.Builder
}
