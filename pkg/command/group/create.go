package group

import (
	"errors"
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type createOptions struct {
	groupLogin string
}

func newCreateCommand(yuqueCli command.Cli) *cobra.Command {
	var opts createOptions

	cmd := &cobra.Command{
		Use:   "create [OPTIONS] NAME",
		Short: "Create a group",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(yuqueCli, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.groupLogin, "group_login", "g", "", "Login of group")

	return cmd
}

func runCreate(yuqueCli command.Cli, name string, opts *createOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	if opts.groupLogin == "" {
		return errors.New("No set group_login")
	}
	c := yuque.NewClient(token)
	_, err := c.Group.Create(name, opts.groupLogin, yuque.GroupOption{})

	return err
}
