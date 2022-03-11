package group

import (
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type UpdateOptions struct {
	name        string
	login       string
	description string
}

func newUpdateCommand(yuqueCli command.Cli) *cobra.Command {
	var opts UpdateOptions

	cmd := &cobra.Command{
		Use:   "update [OPTIONS] GROUPLOGIN",
		Short: "Update group that user join in",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate(yuqueCli, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.name, "name", "n", "", "Name of group")
	flags.StringVarP(&opts.login, "login", "l", "", "Login of group")
	flags.StringVarP(&opts.description, "description", "d", "", "Description of group")

	return cmd
}

func runUpdate(yuqueCli command.Cli, groupLogin string, opts *UpdateOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	_, err := yuque.NewClient(token).Group.Update(groupLogin, yuque.GroupOption{
		Name:        opts.name,
		Login:       opts.login,
		Description: opts.description,
	})

	return err
}
