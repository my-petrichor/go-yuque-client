package repo

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [OPTIONS] NAMESPACE",
		Short: "Get repo info",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(client, args[0])
		},
	}

	return cmd
}

func runGet(client *internal.Client, namespace string) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	_, err = c.Repo.GetInfo(namespace)

	return err
}
