package group

import (
	"fmt"
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [OPTIONS] LOGIN",
		Short: "List all group that user join in",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(args[0])
		},
	}

	return cmd
}

func runList(login string) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	c := yuque.NewClient(token)
	groups, err := c.Group.ListAll(login)
	if err != nil {
		return err
	}

	for _, g := range groups.Data {
		fmt.Printf("group name: %s\ngroup login: %s\n", g.Name, g.Login)
	}

	return nil
}
