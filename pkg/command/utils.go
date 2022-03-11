package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ShowHelp shows the command help.
func ShowHelp() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}

func Login() bool {
	return viper.GetString("token") != ""
}
