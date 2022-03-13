package group

import (
	"os"
	"text/template"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetMemberCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "get_member GROUPLOGIN",
		Short:            "Get a group member info",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetMember(client, args[0])
		},
	}

	return cmd
}

func runGetMember(client *internal.Client, groupLogin string) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	groups, err := c.Group.GetMember(groupLogin)
	if err != nil {
		return err
	}

	var t = `
{{.ID}}.	
login:        {{.Login}}
name:         {{.Name}}
description:  {{.Description}}
`
	groupInfoTemplate, err := template.New("t").Parse(t)
	if err != nil {
		return err
	}

	for i, v := range groups.Data {
		groupInfo := struct {
			ID          int
			Login       string
			Name        string
			Description string
		}{
			ID:          i + 1,
			Login:       v.User.Login,
			Name:        v.User.Name,
			Description: v.User.Description,
		}
		if err = groupInfoTemplate.Execute(os.Stdout, groupInfo); err != nil {
			return err
		}
	}

	return nil
}
