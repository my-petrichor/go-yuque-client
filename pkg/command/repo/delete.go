package repo

import (
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type deleteOptions struct {
	namespace string
}

func newDeleteCommand() *cobra.Command {
	var opts deleteOptions

	cmd := &cobra.Command{
		Use:   "delete [OPTIONS]",
		Short: "Delete a repo (must set namespace flag)",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(&opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.namespace, "namespace", "n", "", "Namespace of repo")

	return cmd
}

func runDelete(opts *deleteOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }

	token := os.Getenv("token")
	_, err := yuque.NewClient(token).Repo.Delete(opts.namespace)

	return err
}
