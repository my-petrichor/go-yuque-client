package group

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newDeleteCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "delete [OPTIONS] GROUPLOGIN",
		Short:            "Delete a group",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(client, args[0])
		},
	}

	return cmd
}

func runDelete(client *internal.Client, groupLogin string) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	_, err = c.Group.Delete(groupLogin)

	return err
}
