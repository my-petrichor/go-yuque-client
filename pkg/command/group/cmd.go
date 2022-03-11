package group

import (
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func NewGroupCommand(yuqueCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group",
		Short: "Manage group",
		Args:  command.NoArgs,
		RunE:  command.ShowHelp(yuqueCli.Err()),
	}
	cmd.AddCommand(
		newListCommand(yuqueCli),
		newListPublicCommand(yuqueCli),
		newGetCommand(yuqueCli),
		newGetMemberCommand(yuqueCli),
		newCreateCommand(yuqueCli),
		newUpdateCommand(yuqueCli),
		newUpdateMemberCommand(yuqueCli),
		newDeleteCommand(yuqueCli),
		newDeleteMemberCommand(yuqueCli),
	)

	return cmd
}
