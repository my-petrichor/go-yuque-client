package repo

import (
	"errors"
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type listOptions struct {
	login       string
	userOrGroup int
}

func newListCommand() *cobra.Command {
	var opts listOptions

	cmd := &cobra.Command{
		Use:   "list [OPTIONS]",
		Short: "List all repo under user or group (must set login flag)",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(&opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.login, "login", "l", "", "Login of user")
	flags.IntVar(&opts.userOrGroup, "user_or_group", 0, "List repo under user or group (0 - user, 1 - group) default 0")

	return cmd
}

func runList(opts *listOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	var err error
	token := os.Getenv("token")
	c := yuque.NewClient(token)
	if opts.userOrGroup == 0 {
		_, err = c.Repo.ListAllUnderUser(opts.login)
	} else if opts.userOrGroup == 1 {
		_, err = c.Repo.ListAllUnderGroup(opts.login)
	} else {
		return errors.New("Error flag userOrGroup")
	}

	return err
}
