package gencmd

import (
	"net/http"

	"github.com/google/wire"

	"github.com/gofunct/hack/pkg/hackcmd"
)

func ProvideGrapiCtx(ctx *Ctx) *hackcmd.Ctx         { return ctx.Ctx }
func ProvideCtx(cmd *Command) *Ctx                   { return cmd.Ctx() }
func ProvideTemplateFS(cmd *Command) http.FileSystem { return cmd.TemplateFS }
func ProvideShouldRun(cmd *Command) ShouldRunFunc    { return cmd.ShouldRun }

// Set contains providers for DI.
var Set = wire.NewSet(
	hackcmd.CtxSet,
	ProvideGrapiCtx,
	ProvideCtx,
	ProvideTemplateFS,
	ProvideShouldRun,
	NewGenerator,
	App{},
)
