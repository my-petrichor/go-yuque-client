package registry

import (
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewLogoutCommand creates a new `yuque logout` command
func NewLogoutCommand(yuqueCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Log out yuque application",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runLogout(yuqueCli)
		},
	}

	return cmd
}

func runLogout(yuqueCli command.Cli) error {
	viper.Set("token", "")

	return viper.WriteConfig()
}
