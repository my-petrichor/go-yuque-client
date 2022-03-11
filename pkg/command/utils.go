package command

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ShowHelp shows the command help.
func ShowHelp(err io.Writer) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetOut(err)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}

func Login() bool {
	return viper.GetString("token") != ""
}
