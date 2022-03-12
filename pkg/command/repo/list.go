package repo

import (
	"errors"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type listOptions struct {
	userOrGroup int
}

func newListCommand(client *internal.Client) *cobra.Command {
	var opts listOptions

	cmd := &cobra.Command{
		Use:   "list [OPTIONS]",
		Short: "List all repo under user or group",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(client, &opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVar(&opts.userOrGroup, "user_or_group", 0, "List repo under user or group (0 - user, 1 - group) default 0")

	return cmd
}

func runList(client *internal.Client, opts *listOptions) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}

	if opts.userOrGroup == 0 {
		_, err = c.Repo.ListAllUnderUser()
	} else if opts.userOrGroup == 1 {
		_, err = c.Repo.ListAllUnderGroup()
	} else {
		return errors.New("Error flag userOrGroup")
	}

	return err
}
