package group

import (
	"errors"
	"os"
	"strconv"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type UpdateMemberOptions struct {
	login      string
	groupLogin string
}

func newUpdateMemberCommand() *cobra.Command {
	var opts UpdateMemberOptions

	cmd := &cobra.Command{
		Use:   "update_member [OPTIONS] ROLE",
		Short: "Update group member authority (0 - manager, 1 - ordinary)",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			role, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			return runUpdateMember(role, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.login, "login", "l", "", "Login of group member")
	flags.StringVarP(&opts.groupLogin, "group_login", "g", "", "Login of group")

	return cmd
}

func runUpdateMember(role int, opts *UpdateMemberOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	if opts.groupLogin == "" {
		return errors.New("No set group_login flag")
	} else if opts.login == "" {
		return errors.New("No set login flag")
	}

	c := yuque.NewClient(token)
	_, err := c.Group.UpdateMember(opts.groupLogin, opts.login, role)

	return err
}
