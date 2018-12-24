//+build wireinject

package svcgen

import (
	"github.com/google/wire"

	"github.com/gofunct/hack/pkg/cli"
	"github.com/gofunct/hack/pkg/gencmd"
	"github.com/gofunct/hack/pkg/protoc"
)

func NewApp(*gencmd.Command) (*App, error) {
	wire.Build(
		Set,
		gencmd.Set,
		cli.UIInstance,
		protoc.WrapperSet,
	)
	return nil, nil
}
