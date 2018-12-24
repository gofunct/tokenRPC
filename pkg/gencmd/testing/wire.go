//+build wireinject

package testing

import (
	"github.com/google/wire"

	"github.com/gofunct/hack/pkg/cli"
	"github.com/gofunct/hack/pkg/gencmd"
)

func NewTestApp(*gencmd.Command, cli.UI) (*gencmd.App, error) {
	wire.Build(
		gencmd.Set,
	)
	return nil, nil
}
