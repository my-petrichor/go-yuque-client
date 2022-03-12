package document

import (
	"fmt"

	huge "github.com/dablelv/go-huge-util"
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type getOptions struct {
}

func newGetCommand(client *internal.Client) *cobra.Command {
	var opts getOptions

	cmd := &cobra.Command{
		Use:              "get [OPTIONS] PATH",
		Short:            "Get user info",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(client, args[0], &opts)
		},
	}

	return cmd
}

func runGet(client *internal.Client, path string, opts *getOptions) error {
	namespace, slug, err := splitPath(path)
	if err != nil {
		return err
	}

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	u, err := c.Document.GetInfo(namespace, slug)
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
