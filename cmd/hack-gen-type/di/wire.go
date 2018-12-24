//+build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/gofunct/hack/pkg/cli"
	"github.com/gofunct/hack/pkg/gencmd"
	"github.com/gofunct/hack/pkg/protoc"
)

func NewApp(*gencmd.Command) (*App, error) {
	wire.Build(
		App{},
		gencmd.Set,
		cli.UIInstance,
		protoc.WrapperSet,
	)
	return nil, nil
}
