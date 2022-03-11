package repo

import (
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type getDirOptions struct {
	namespace string
}

func newGetDirCommand(yuqueCli command.Cli) *cobra.Command {
	var opts getDirOptions

	cmd := &cobra.Command{
		Use:   "get_dir [OPTIONS]",
		Short: "Get repo dir (must set namespace flag)",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetDir(yuqueCli, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.namespace, "namespace", "n", "", "Namespace of repo")

	return cmd
}

func runGetDir(yuqueCli command.Cli, opts *getDirOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }

	token := os.Getenv("token")
	_, err := yuque.NewClient(token).Repo.GetDir(opts.namespace)

	return err
}
