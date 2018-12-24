package gencmd

import (
	"github.com/gofunct/hack/pkg/hackcmd"
)

// Option configures a command context.
type Option func(*Ctx)

// WithGrapiCtx specifies a hack command context.
func WithGrapiCtx(gctx *hackcmd.Ctx) Option {
	return func(ctx *Ctx) {
		ctx.Ctx = gctx
	}
}

// WithCreateAppFunc specifies a dependencies initializer.
func WithCreateAppFunc(f CreateAppFunc) Option {
	return func(ctx *Ctx) {
		ctx.CreateAppFunc = f
	}
}
