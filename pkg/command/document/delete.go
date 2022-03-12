package document

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newDeleteCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [OPTIONS] PATH",
		Short: "Delete a document",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(client, args[0])
		},
	}

	return cmd
}

func runDelete(client *internal.Client, path string) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	namespace, slug, err := splitPath(path)
	if err != nil {
		return err
	}

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	_, err = c.Document.Delete(namespace, slug)

	return err
}
