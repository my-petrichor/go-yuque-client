package user

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get user info",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(client)
		},
	}

	return cmd
}

func runGet(client *internal.Client) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	u, err := yuque.NewClient(client.Token).User.GetInfo()
	if err != nil {
		return err
	}

	fmt.Printf("name: %s\nlogin: %s\n", u.Data.Name, u.Data.Login)

	return nil
}
