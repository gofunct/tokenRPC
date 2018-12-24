// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package svcgen

import (
	"github.com/gofunct/hack/pkg/cli"
	"github.com/gofunct/hack/pkg/gencmd"
	"github.com/gofunct/hack/pkg/hackcmd"
	"github.com/gofunct/hack/pkg/protoc"
)

// Injectors from wire.go:

func NewApp(command *gencmd.Command) (*App, error) {
	ctx := gencmd.ProvideCtx(command)
	hackcmdCtx := gencmd.ProvideGrapiCtx(ctx)
	config := hackcmd.ProvideProtocConfig(hackcmdCtx)
	fs := hackcmd.ProvideFS(hackcmdCtx)
	execInterface := hackcmd.ProvideExecer(hackcmdCtx)
	io := hackcmd.ProvideIO(hackcmdCtx)
	ui := cli.UIInstance(io)
	rootDir := hackcmd.ProvideRootDir(hackcmdCtx)
	toolsConfig := protoc.ProvideGexConfig(fs, execInterface, io, rootDir)
	repository, err := protoc.ProvideToolRepository(toolsConfig)
	if err != nil {
		return nil, err
	}
	wrapper := protoc.NewWrapper(config, fs, execInterface, ui, repository, rootDir)
	hackcmdConfig := hackcmd.ProvideConfig(hackcmdCtx)
	builder := ProvideParamsBuilder(rootDir, config, hackcmdConfig)
	app := &App{
		ProtocWrapper: wrapper,
		ParamsBuilder: builder,
	}
	return app, nil
}
