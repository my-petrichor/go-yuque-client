package group

import (
	"errors"
	"strconv"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type UpdateMemberOptions struct {
	groupLogin string
}

func newUpdateMemberCommand(client *internal.Client) *cobra.Command {
	var opts UpdateMemberOptions

	cmd := &cobra.Command{
		Use:   "update_member [OPTIONS] ROLE",
		Short: "Update group member authority (0 - manager, 1 - ordinary), (must set group_login flag)",
		Args:  command.ExactArgs(1),
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

	return cmd
}

func runUpdateMember(client *internal.Client, role int, opts *UpdateMemberOptions) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}
	if opts.groupLogin == "" {
		return errors.New("No set group_login flag")
	} else if opts.groupLogin == "" {
		return errors.New("No set group_login flag")
	}

	c := yuque.NewClient(client.Token)
	_, err := c.Group.UpdateMember(opts.groupLogin, role)

	return err
}
