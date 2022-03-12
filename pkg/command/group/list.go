package group

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newListCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [OPTIONS]",
		Short: "List all group that user join in",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(client)
		},
	}

	return cmd
}

func runList(client *internal.Client) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	groups, err := yuque.NewClient(client.Token).Group.ListAll()
	if err != nil {
		return err
	}

	for _, g := range groups.Data {
		fmt.Printf("group name: %s\ngroup login: %s\n", g.Name, g.Login)
	}

	return nil
}
