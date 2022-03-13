package group

import (
	"os"
	"text/template"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newListCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "list [OPTIONS]",
		Short:            "List all group that user join in",
		Args:             command.NoArgs,
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(client)
		},
	}

	return cmd
}

func runList(client *internal.Client) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	groups, err := c.Group.ListAll()
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
			Login:       v.Login,
			Name:        v.Name,
			Description: v.Description,
		}
		if err = groupInfoTemplate.Execute(os.Stdout, groupInfo); err != nil {
			return err
		}
	}

	return nil
}
