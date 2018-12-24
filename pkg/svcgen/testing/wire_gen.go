// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package testing

import (
	"github.com/gofunct/hack/pkg/cli"
	"github.com/gofunct/hack/pkg/gencmd"
	"github.com/gofunct/hack/pkg/hackcmd"
	"github.com/gofunct/hack/pkg/protoc"
	"github.com/gofunct/hack/pkg/svcgen"
)

// Injectors from wire.go:

func NewTestApp(command *gencmd.Command, wrapper protoc.Wrapper, ui cli.UI) (*svcgen.App, error) {
	ctx := gencmd.ProvideCtx(command)
	hackcmdCtx := gencmd.ProvideGrapiCtx(ctx)
	rootDir := hackcmd.ProvideRootDir(hackcmdCtx)
	config := hackcmd.ProvideProtocConfig(hackcmdCtx)
	hackcmdConfig := hackcmd.ProvideConfig(hackcmdCtx)
	builder := svcgen.ProvideParamsBuilder(rootDir, config, hackcmdConfig)
	app := &svcgen.App{
		ProtocWrapper: wrapper,
		ParamsBuilder: builder,
	}
	return app, nil
}
