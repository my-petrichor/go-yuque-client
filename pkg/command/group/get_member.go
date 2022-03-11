package group

import (
	"fmt"
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetMemberCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get_member [OPTIONS] GROUPLOGIN",
		Short: "Get a group member info",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetMember(args[0])
		},
	}

	return cmd
}

func runGetMember(groupLogin string) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	c := yuque.NewClient(token)
	g, err := c.Group.GetMembers(groupLogin)
	if err != nil {
		return err
	}
	for _, m := range g.Data {
		fmt.Printf("user name: %s\n", m.User.Name)
	}

	return nil
}
