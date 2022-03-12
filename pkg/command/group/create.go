package group

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type createOptions struct {
	groupName   string
	description string
}

func newCreateCommand(client *internal.Client) *cobra.Command {
	var opts createOptions

	cmd := &cobra.Command{
		Use:   "create [OPTIONS] GROUPLOGIN",
		Short: "Create a group",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(client, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.groupName, "name", "n", "", "Name of group")
	flags.StringVarP(&opts.description, "description", "d", "", "Description of group")

	return cmd
}

func runCreate(client *internal.Client, groupLogin string, opts *createOptions) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	_, err = c.Group.Create(groupLogin, opts.groupName, yuque.GroupOption{
		Description: opts.description,
	})

	return err
}
