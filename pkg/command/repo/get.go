package repo

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
		Use:              "get [OPTIONS] NAMESPACE",
		Short:            "Get repo info",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(client, args[0])
		},
	}

	return cmd
}

func runGet(client *internal.Client, namespace string) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	r, err := c.Repo.GetInfo(namespace)
	if err != nil {
		return err
	}

	data, err := huge.ToIndentJSON(&r.Data)
	if err != nil {
		return err
	}
	fmt.Println(data)

	return nil
}
