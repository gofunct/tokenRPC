package svcgen

import (
	"github.com/google/wire"

	"github.com/gofunct/hack/pkg/cli"
	"github.com/gofunct/hack/pkg/hackcmd"
	"github.com/gofunct/hack/pkg/protoc"
	"github.com/gofunct/hack/pkg/svcgen/params"
)

func ProvideParamsBuilder(rootDir cli.RootDir, protocCfg *protoc.Config, hackCfg *hackcmd.Config) params.Builder {
	return params.NewBuilder(
		rootDir,
		protocCfg.ProtosDir,
		protocCfg.OutDir,
		hackCfg.Hack.ServerDir,
		hackCfg.Package,
	)
}

var Set = wire.NewSet(
	ProvideParamsBuilder,
	App{},
)