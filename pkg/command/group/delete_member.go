package group

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type deleteMemberOptions struct {
	groupLogin string
}

func newDeleteMemberCommand(client *internal.Client) *cobra.Command {
	var opts deleteMemberOptions

	cmd := &cobra.Command{
		Use:              "delete_member LOGIN",
		Short:            "Delete a group member",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDeleteMember(client, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.groupLogin, "group_login", "g", "", "Login of group")

	cmd.MarkFlagRequired("group_login")

	return cmd
}

func runDeleteMember(client *internal.Client, login string, opts *deleteMemberOptions) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	_, err = c.Group.DeleteMember(opts.groupLogin, login)

	return err
}
