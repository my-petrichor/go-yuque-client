package repo

import (
	"fmt"
	"os"
	"text/template"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type listOptions struct {
	groupLogin string
}

func newListCommand(client *internal.Client) *cobra.Command {
	var opts listOptions

	cmd := &cobra.Command{
		Use:              "list [OPTIONS]",
		Short:            "List all repo under user or group",
		Args:             command.NoArgs,
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(client, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.groupLogin, "group_login", "g", "", "List repo under group")

	return cmd
}

func runList(client *internal.Client, opts *listOptions) error {
	var (
		repos *yuque.ResponseBookSerializer
		err   error
	)

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}

	if opts.groupLogin == "" {
		repos, err = c.Repo.ListAllUnderUser()
	} else {
		fmt.Println("----")
		// r, err = c.Repo.ListAllUnderGroup(opts.groupLogin)
	}
	if err != nil {
		return err
	}

	var t = `
{{.ID}}.
name:       {{.Name}}
namespace:  {{.Namespace}}
		`

	repoInfoTemplate, err := template.New("t").Parse(t)
	if err != nil {
		return err
	}

	for i, v := range repos.Data {
		repoInfo := struct {
			ID        int
			Name      string
			Namespace string
		}{
			ID:        i + 1,
			Name:      v.Name,
			Namespace: v.Namespace,
		}
		if err = repoInfoTemplate.Execute(os.Stdout, repoInfo); err != nil {
			return err
		}
	}

	return nil
}
