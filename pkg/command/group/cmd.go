package group

import (
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func NewGroupCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group",
		Short: "Manage group",
		Args:  command.NoArgs,
		RunE:  command.ShowHelp(),
	}
	cmd.AddCommand(
		newListCommand(),
		newListPublicCommand(),
		newGetCommand(),
		newGetMemberCommand(),
		newCreateCommand(),
		newUpdateCommand(),
		newUpdateMemberCommand(),
		newDeleteCommand(),
		newDeleteMemberCommand(),
	)

	return cmd
}
