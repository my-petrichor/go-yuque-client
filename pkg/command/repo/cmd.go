package repo

import (
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func NewRepoCommand(yuqueCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Manage repo",
		Args:  command.NoArgs,
		RunE:  command.ShowHelp(yuqueCli.Err()),
	}
	cmd.AddCommand(
		newListCommand(yuqueCli),
		newGetCommand(yuqueCli),
		newGetDirCommand(yuqueCli),
		newCreateCommand(yuqueCli),
		newUpdateCommand(yuqueCli),
		newDeleteCommand(yuqueCli),
	)

	return cmd
}
