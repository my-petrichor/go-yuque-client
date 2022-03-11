package document

import (
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newDeleteCommand(yuqueCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [OPTIONS] PATH",
		Short: "Delete a document",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(yuqueCli, args[0])
		},
	}

	return cmd
}

func runDelete(yuqueCli command.Cli, path string) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	namespace, slug, err := splitPath(path)
	if err != nil {
		return err
	}
	token := os.Getenv("token")
	_, err = yuque.NewClient(token).Document.Delete(namespace, slug)

	return err
}
