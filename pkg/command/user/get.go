package user

import (
	"fmt"

	huge "github.com/dablelv/go-huge-util"
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "get",
		Short:            "Get user info",
		Args:             command.NoArgs,
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(client)
		},
	}

	return cmd
}

func runGet(client *internal.Client) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}

	u, err := c.User.GetInfo()
	if err != nil {
		return err
	}

	data, err := huge.ToIndentJSON(&u.Data)
	if err != nil {
		return err
	}
	fmt.Println(data)

	return nil
}
