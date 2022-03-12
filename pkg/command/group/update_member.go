package group

import (
	"strconv"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type UpdateMemberOptions struct {
	groupLogin string
	login      string
}

func newUpdateMemberCommand(client *internal.Client) *cobra.Command {
	var opts UpdateMemberOptions

	cmd := &cobra.Command{
		Use:              "update_member [OPTIONS] ROLE",
		Short:            "Update group member authority (0 - manager, 1 - ordinary)",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			role, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return runUpdateMember(client, role, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.groupLogin, "group_login", "g", "", "Login of group")
	flags.StringVarP(&opts.login, "login", "l", "", "Login of user who need to update")

	cmd.MarkFlagRequired("group_login")
	cmd.MarkFlagRequired("login")

	return cmd
}

func runUpdateMember(client *internal.Client, role int, opts *UpdateMemberOptions) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	_, err = c.Group.UpdateMember(opts.groupLogin, opts.login, role)

	return err
}
