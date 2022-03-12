package group

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get GROUPLOGIN",
		Short: "Get a group info",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetMember(client, args[0])
		},
	}

	return cmd
}

func runGet(client *internal.Client, groupLogin string) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	g, err := c.Group.GetInfo(groupLogin)
	if err != nil {
		return err
	}
	fmt.Printf("user name: %s\nuser login: %s\n", g.Data.Name, g.Data.Login)

	return nil
}
