package document

import (
	"fmt"
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type getOptions struct {
	name bool
}

func newGetCommand(yuqueCli command.Cli) *cobra.Command {
	var opts getOptions

	cmd := &cobra.Command{
		Use:   "get [OPTIONS]",
		Short: "Get user info",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(yuqueCli, &opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.name, "name", "n", false, "Only display user name")

	return cmd
}

func runGet(yuqueCli command.Cli, opts *getOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	c := yuque.NewClient(token)
	u, err := c.User.GetInfo()
	if err != nil {
		return err
	}
	if opts.name {
		fmt.Printf("name: %s\n", u.Data.Name)
	} else {
		fmt.Printf("name: %s\nlogin: %s\n", u.Data.Name, u.Data.Login)
	}

	return nil
}
