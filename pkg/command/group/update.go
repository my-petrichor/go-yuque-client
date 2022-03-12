package group

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type UpdateOptions struct {
	name        string
	login       string
	description string
}

func newUpdateCommand(client *internal.Client) *cobra.Command {
	var opts UpdateOptions

	cmd := &cobra.Command{
		Use:   "update [OPTIONS] GROUPLOGIN",
		Short: "Update group that user join in",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate(client, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.name, "name", "n", "", "Name of group")
	flags.StringVarP(&opts.login, "login", "l", "", "Login of group")
	flags.StringVarP(&opts.description, "description", "d", "", "Description of group")

	return cmd
}

func runUpdate(client *internal.Client, groupLogin string, opts *UpdateOptions) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}
	_, err := yuque.NewClient(client.Token).Group.Update(groupLogin, yuque.GroupOption{
		Name:        opts.name,
		Login:       opts.login,
		Description: opts.description,
	})

	return err
}
