//+build wireinject

package testing

import (
	"github.com/google/wire"

	"github.com/gofunct/hack/pkg/cli"
	"github.com/gofunct/hack/pkg/gencmd"
	"github.com/gofunct/hack/pkg/protoc"
	"github.com/gofunct/hack/pkg/svcgen"
)

func NewTestApp(*gencmd.Command, protoc.Wrapper, cli.UI) (*svcgen.App, error) {
	wire.Build(
		gencmd.Set,
		svcgen.ProvideParamsBuilder,
		svcgen.App{},
	)
	return nil, nil
}
