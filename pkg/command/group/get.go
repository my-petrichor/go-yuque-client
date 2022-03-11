package group

import (
	"fmt"
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetCommand(yuqueCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [OPTIONS] GROUPLOGIN",
		Short: "Get a group info",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetMember(yuqueCli, args[0])
		},
	}

	return cmd
}

func runGet(yuqueCli command.Cli, groupLogin string) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	c := yuque.NewClient(token)
	g, err := c.Group.GetInfo(groupLogin)
	if err != nil {
		return err
	}
	fmt.Printf("user name: %s\nuser login: %s\n", g.Data.Name, g.Data.Login)

	return nil
}
