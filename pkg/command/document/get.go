package document

import (
	"fmt"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type getOptions struct {
	name bool
}

func newGetCommand(client *internal.Client) *cobra.Command {
	var opts getOptions

	cmd := &cobra.Command{
		Use:   "get [OPTIONS]",
		Short: "Get user info",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(client, &opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.name, "name", "n", false, "Only display user name")

	return cmd
}

func runGet(client *internal.Client, opts *getOptions) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	u, err := yuque.NewClient(client.Token).User.GetInfo()
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
