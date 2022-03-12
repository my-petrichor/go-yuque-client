package repo

import (
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func NewRepoCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Manage repo",
		Args:  command.NoArgs,
		RunE:  command.ShowHelp(),
	}
	cmd.AddCommand(
		newListCommand(client),
		newGetCommand(client),
		newGetDirCommand(client),
		newCreateCommand(client),
		newUpdateCommand(client),
		newDeleteCommand(client),
	)

	return cmd
}
