package di

import (
	"github.com/gofunct/hack/pkg/gencmd"
	"github.com/gofunct/hack/pkg/protoc"
)

type CreateAppFunc func(*gencmd.Command) (*App, error)

type App struct {
	Protoc protoc.Wrapper
}
