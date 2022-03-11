package user

import (
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func NewUserCommand(yuqueCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage user",
		Args:  command.NoArgs,
		RunE:  command.ShowHelp(yuqueCli.Err()),
	}
	cmd.AddCommand(
		newGetCommand(yuqueCli),
	)

	return cmd
}
