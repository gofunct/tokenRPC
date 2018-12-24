//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/gofunct/hack/pkg/tool"

	"github.com/gofunct/hack/pkg/cli"
	"github.com/gofunct/hack/pkg/excmd"
	"github.com/gofunct/hack/pkg/hackcmd"
	"github.com/gofunct/hack/pkg/hackcmd/internal/module"
	"github.com/gofunct/hack/pkg/hackcmd/internal/usecase"
	"github.com/gofunct/hack/pkg/protoc"
)

func NewUI(*hackcmd.Ctx) cli.UI {
	wire.Build(Set)
	return nil
}

func NewCommandExecutor(*hackcmd.Ctx) excmd.Executor {
	wire.Build(Set)
	return nil
}

func NewGenerator(*hackcmd.Ctx) module.Generator {
	wire.Build(Set)
	return nil
}

func NewScriptLoader(*hackcmd.Ctx) module.ScriptLoader {
	wire.Build(Set)
	return nil
}

func NewToolRepository(*hackcmd.Ctx) (tool.Repository, error) {
	wire.Build(Set)
	return nil, nil
}

func NewProtocWrapper(*hackcmd.Ctx) (protoc.Wrapper, error) {
	wire.Build(Set)
	return nil, nil
}

func NewInitializeProjectUsecase(*hackcmd.Ctx) usecase.InitializeProjectUsecase {
	wire.Build(Set)
	return nil
}
