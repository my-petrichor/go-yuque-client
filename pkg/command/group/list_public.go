package group

import (
	"fmt"

	huge "github.com/dablelv/go-huge-util"
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newListPublicCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "list_public [OPTIONS]",
		Short:            "List all public group",
		Args:             command.NoArgs,
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListPublic(client)
		},
	}

	return cmd
}

func runListPublic(client *internal.Client) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	groups, err := c.Group.ListPublic()
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
