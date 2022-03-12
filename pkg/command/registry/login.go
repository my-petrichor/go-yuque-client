package registry

import (
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

// NewLoginCommand creates a new `yuque login` command
func NewLoginCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login TOKEN",
		Short: "Log in yuque application",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runLogin(client, args[0])
		},
	}

	return cmd
}

func runLogin(client *internal.Client, token string) error {
	return client.Login(token)
}
