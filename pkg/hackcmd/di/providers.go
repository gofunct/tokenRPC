package di

import (
	"github.com/google/wire"
	"github.com/gofunct/hack/tools"

	"github.com/gofunct/hack/pkg/cli"
	"github.com/gofunct/hack/pkg/excmd"
	"github.com/gofunct/hack/pkg/hackcmd"
	"github.com/gofunct/hack/pkg/hackcmd/internal/module"
	"github.com/gofunct/hack/pkg/hackcmd/internal/module/generator"
	"github.com/gofunct/hack/pkg/hackcmd/internal/module/script"
	"github.com/gofunct/hack/pkg/hackcmd/internal/usecase"
	"github.com/gofunct/hack/pkg/protoc"
)

func ProvideGenerator(ctx *hackcmd.Ctx, ui cli.UI) module.Generator {
	return generator.New(
		ctx.FS,
		ui,
	)
}

func ProvideScriptLoader(ctx *hackcmd.Ctx, executor excmd.Executor) module.ScriptLoader {
	return script.NewLoader(ctx.FS, executor, ctx.RootDir.String())
}

func ProvideInitializeProjectUsecase(ctx *hackcmd.Ctx, gexCfg *tools.Config, ui cli.UI, generator module.Generator) usecase.InitializeProjectUsecase {
	return usecase.NewInitializeProjectUsecase(
		ui,
		generator,
		gexCfg,
	)
}

var Set = wire.NewSet(
	hackcmd.CtxSet,
	protoc.WrapperSet,
	cli.UIInstance,
	excmd.NewExecutor,
	ProvideGenerator,
	ProvideScriptLoader,
	ProvideInitializeProjectUsecase,
)
