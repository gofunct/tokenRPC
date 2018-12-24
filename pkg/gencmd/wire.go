//+build wireinject

package gencmd

import (
	"github.com/google/wire"
	"github.com/gofunct/hack/pkg/cli"
)

func newApp(*Command) (*App, error) {
	wire.Build(
		Set,
		cli.UIInstance,
	)
	return nil, nil
}
