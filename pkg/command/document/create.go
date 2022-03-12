package document

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type createOptions struct {
	slug   string
	title  string
	body   string
	format string
	public int
}

func newCreateCommand(client *internal.Client) *cobra.Command {
	var opts createOptions

	cmd := &cobra.Command{
		Use:              "create [OPTIONS] NAMESPACE",
		Short:            "Create a document",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(client, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.slug, "slug", "s", "", "Slug of document")
	flags.StringVarP(&opts.title, "title", "t", "", "Title of document")
	flags.StringVarP(&opts.body, "content", "c", "", "Content of document")
	flags.StringVarP(&opts.format, "format", "f", "markdown", "Type of document (default markdown)")
	flags.IntVarP(&opts.public, "public", "p", 0, "Public of document (0 - private, 1 - public)")

	return cmd
}

func runCreate(client *internal.Client, namespace string, opts *createOptions) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	_, err = c.Document.Create(namespace, yuque.DocumentOption{
		Format: opts.format,
		Public: opts.public,
		Body:   opts.body,
		Title:  opts.title,
		Slug:   opts.slug,
	})

	return err
}
