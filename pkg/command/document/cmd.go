package document

import (
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func NewDocumentCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "document",
		Short: "Manage document",
		Args:  command.NoArgs,
		RunE:  command.ShowHelp(),
	}
	cmd.AddCommand(
		newListCommand(client),
		newGetCommand(client),
		newCreateCommand(client),
		newUpdateCommand(client),
		newDeleteCommand(client),
	)

	return cmd
}
