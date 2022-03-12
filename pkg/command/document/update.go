package document

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type updateOptions struct {
	slug     string
	title    string
	body     string
	format   string
	public   int
	forceASL int
}

func newUpdateCommand(client *internal.Client) *cobra.Command {
	var opts updateOptions

	cmd := &cobra.Command{
		Use:   "update [OPTIONS] PATH",
		Short: "Update a document",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate(client, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.slug, "slug", "s", "", "Update document slug")
	flags.StringVarP(&opts.title, "title", "t", "", "Update document title")
	flags.StringVarP(&opts.body, "body", "b", "", "Update document body")
	flags.StringVarP(&opts.format, "format", "f", "markdown", "Update document type")
	flags.IntVar(&opts.forceASL, "force_asl", 0, "force_asl = 1 ensure the correct conversion of the content (default 0)")
	flags.IntVarP(&opts.public, "public", "p", 0, "Update document public (0 - private, 1 - public)")

	return cmd
}

func runUpdate(client *internal.Client, path string, opts *updateOptions) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	namespace, slug, err := splitPath(path)
	if err != nil {
		return err
	}

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	_, err = c.Document.Update(namespace, slug, yuque.DocumentOption{
		Format:   opts.format,
		Public:   opts.public,
		Body:     opts.body,
		Title:    opts.title,
		Slug:     opts.slug,
		ForceASL: opts.forceASL,
	})

	return err
}
