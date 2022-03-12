package group

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetMemberCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get_member GROUPLOGIN",
		Short: "Get a group member info",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetMember(client, args[0])
		},
	}

	return cmd
}

func runGetMember(client *internal.Client, groupLogin string) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}
	g, err := yuque.NewClient(client.Token).Group.GetMember(groupLogin)
	if err != nil {
		return err
	}
	for _, m := range g.Data {
		fmt.Printf("user name: %s\n", m.User.Name)
	}

	return nil
}
