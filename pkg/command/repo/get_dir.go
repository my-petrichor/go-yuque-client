package repo

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetDirCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get_dir [OPTIONS] NAMESPACE",
		Short: "Get repo dir",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetDir(client, args[0])
		},
	}

	return cmd
}

func runGetDir(client *internal.Client, namespace string) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	_, err = c.Repo.GetDir(namespace)

	return err
}
