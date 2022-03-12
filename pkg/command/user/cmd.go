package user

import (
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func NewUserCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage user",
		Args:  command.NoArgs,
		RunE:  command.ShowHelp(),
	}
	cmd.AddCommand(
		newGetCommand(client),
	)

	return cmd
}
