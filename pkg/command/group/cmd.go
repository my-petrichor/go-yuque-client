package group

import (
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func NewGroupCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group",
		Short: "Manage group",
		Args:  command.NoArgs,
		RunE:  command.ShowHelp(),
	}
	cmd.AddCommand(
		newListCommand(client),
		newListPublicCommand(client),
		newGetCommand(client),
		newGetMemberCommand(client),
		newCreateCommand(client),
		newUpdateCommand(client),
		newUpdateMemberCommand(client),
		newDeleteCommand(client),
		newDeleteMemberCommand(client),
	)

	return cmd
}
