package document

import (
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func NewDocumentCommand(yuqueCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "document",
		Short: "Manage document",
		Args:  command.NoArgs,
		RunE:  command.ShowHelp(yuqueCli.Err()),
	}
	cmd.AddCommand(
		newListCommand(yuqueCli),
		newGetCommand(yuqueCli),
		newCreateCommand(yuqueCli),
		newUpdateCommand(yuqueCli),
		newDeleteCommand(yuqueCli),
	)

	return cmd
}
