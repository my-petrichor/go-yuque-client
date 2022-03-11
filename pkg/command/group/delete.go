package group

import (
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newDeleteCommand(yuqueCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [OPTIONS] GROUPLOGIN",
		Short: "Delete a group",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(yuqueCli, args[0])
		},
	}

	return cmd
}

func runDelete(yuqueCli command.Cli, groupName string) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	c := yuque.NewClient(token)
	_, err := c.Group.Delete(groupName)

	return err
}
