package registry

import (
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

// NewLogoutCommand creates a new `yuque logout` command
func NewLogoutCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Log out yuque application",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runLogout(client)
		},
	}

	return cmd
}

func runLogout(client *internal.Client) error {
	return client.Logout()
}
