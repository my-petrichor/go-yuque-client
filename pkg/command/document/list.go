package document

import (
	"os"
	"text/template"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type listOptions struct {
	slug   string
	title  string
	body   string
	format string
	public int
}

func newListCommand(client *internal.Client) *cobra.Command {
	var opts listOptions

	cmd := &cobra.Command{
		Use:              "list [OPTIONS] NAMESPACE",
		Short:            "List all document under a repo",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(client, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.slug, "slug", "s", "", "Document of slug")
	flags.StringVarP(&opts.title, "title", "t", "", "Only display user name")
	flags.StringVarP(&opts.body, "body", "b", "", "Only display user name")
	flags.StringVarP(&opts.format, "format", "f", "markdown", "Type of document (default markdown)")
	flags.IntVarP(&opts.public, "public", "p", 0, "Only display user name")

	return cmd
}

func runList(client *internal.Client, namespace string, opts *listOptions) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	documents, err := c.Document.ListAll(namespace)
	if err != nil {
		return err
	}

	var t = `
{{.ID}}.
title:  {{.Title}}
slug:   {{.Slug}}
   		`

	docInfoTemplate, err := template.New("t").Parse(t)
	if err != nil {
		return err
	}

	for i, doc := range documents.Data {
		docInfo := struct {
			ID    int
			Title string
			Slug  string
		}{
			ID:    i + 1,
			Title: doc.Title,
			Slug:  doc.Slug,
		}
		if err = docInfoTemplate.Execute(os.Stdout, docInfo); err != nil {
			return err
		}
	}

	return nil
}
