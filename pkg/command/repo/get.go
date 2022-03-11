package repo

import (
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type getOptions struct {
	namespace string
}

func newGetCommand(yuqueCli command.Cli) *cobra.Command {
	var opts getOptions

	cmd := &cobra.Command{
		Use:   "get [OPTIONS]",
		Short: "Get repo info (must set namespace flag)",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(yuqueCli, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.namespace, "namespace", "n", "", "Namespace of repo")

	return cmd
}

func runGet(yuqueCli command.Cli, opts *getOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }

	token := os.Getenv("token")
	_, err := yuque.NewClient(token).Repo.GetInfo(opts.namespace)

	return err
}
