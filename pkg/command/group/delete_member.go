package group

import (
	"errors"
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type deleteMemberOptions struct {
	groupLogin string
}

func newDeleteMemberCommand(yuqueCli command.Cli) *cobra.Command {
	var opts createOptions

	cmd := &cobra.Command{
		Use:   "delete_member [OPTIONS] LOGIN",
		Short: "Delete a group member",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(yuqueCli, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.groupLogin, "group_login", "g", "", "Login of group")

	return cmd
}

func runDeleteMember(yuqueCli command.Cli, login string, opts *createOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	if opts.groupLogin == "" {
		return errors.New("No set group_login")
	}
	c := yuque.NewClient(token)
	_, err := c.Group.DeleteMember(login, opts.groupLogin)

	return err
}
