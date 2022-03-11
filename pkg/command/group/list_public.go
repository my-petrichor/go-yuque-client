package group

import (
	"fmt"
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newListPublicCommand(yuqueCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list_public [OPTIONS]",
		Short: "List all public group",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListPublic(yuqueCli)
		},
	}

	return cmd
}

func runListPublic(yuqueCli command.Cli) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	c := yuque.NewClient(token)
	groups, err := c.Group.ListPublic()
	if err != nil {
		return err
	}

	for _, g := range groups.Data {
		fmt.Printf("group name: %s\ngroup login: %s\n", g.Name, g.Login)
	}

	return nil
}
