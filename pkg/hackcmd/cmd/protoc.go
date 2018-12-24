package cmd

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/gofunct/hack/pkg/hackcmd"
	"github.com/gofunct/hack/pkg/hackcmd/di"
)

func newProtocCommand(ctx *hackcmd.Ctx) *cobra.Command {
	return &cobra.Command{
		Use:           "protoc",
		Short:         "Run protoc",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !ctx.IsInsideApp() {
				return errors.New("protoc command should be execute inside a hack application directory")
			}
			protocw, err := di.NewProtocWrapper(ctx)
			if err != nil {
				return errors.WithStack(err)
			}
			return errors.WithStack(protocw.Exec(context.TODO()))
		},
	}
}
