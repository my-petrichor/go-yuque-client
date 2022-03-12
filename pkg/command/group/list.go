package group

import (
	"fmt"

	huge "github.com/dablelv/go-huge-util"
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newListCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "list [OPTIONS]",
		Short:            "List all group that user join in",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(client)
		},
	}

	return cmd
}

func runList(client *internal.Client) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	groups, err := c.Group.ListAll()
	if err != nil {
		return err
	}

	data, err := huge.ToIndentJSON(&groups.Data)
	if err != nil {
		return err
	}
	fmt.Println(data)

	return nil
}
