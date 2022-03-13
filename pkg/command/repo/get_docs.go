package repo

import (
	"os"
	"text/template"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetDirCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "get_docs [OPTIONS] NAMESPACE",
		Short:            "Get documents info under the repo",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetDir(client, args[0])
		},
	}

	return cmd
}

func runGetDir(client *internal.Client, namespace string) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	documents, err := c.Repo.GetDir(namespace)
	if err != nil {
		return err
	}

	var t = `
{{.ID}}.
title: {{.Title}}
type:  {{.Type}}
slug:  {{.Slug}}
`

	for i, doc := range documents.Data {
		docInfo := struct {
			ID    int
			Title string
			Type  string
			Slug  string
		}{
			ID:    i + 1,
			Title: doc.Title,
			Type:  doc.Type,
			Slug:  doc.Slug,
		}
		repoInfoTemplate, err := template.New("t").Parse(t)
		if err != nil {
			return err
		}
		if err = repoInfoTemplate.Execute(os.Stdout, docInfo); err != nil {
			return err
		}
	}

	return nil
}
