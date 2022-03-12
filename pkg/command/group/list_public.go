package group

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newListPublicCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list_public [OPTIONS]",
		Short: "List all public group",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListPublic(client)
		},
	}

	return cmd
}

func runListPublic(client *internal.Client) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	groups, err := yuque.NewClient(client.Token).Group.ListPublic()
	if err != nil {
		return err
	}

	for _, g := range groups.Data {
		fmt.Printf("group name: %s\ngroup login: %s\n", g.Name, g.Login)
	}

	return nil
}
