package group

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newDeleteMemberCommand(client *internal.Client) *cobra.Command {
	var opts createOptions

	cmd := &cobra.Command{
		Use:   "delete_member GROUPLOGIN",
		Short: "Delete a group member",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(client, args[0], &opts)
		},
	}

	return cmd
}

func runDeleteMember(client *internal.Client, groupLogin string) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}
	_, err := yuque.NewClient(client.Token).Group.DeleteMember(groupLogin)

	return err
}
