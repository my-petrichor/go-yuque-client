package document

import (
	"fmt"
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
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

func newListCommand() *cobra.Command {
	var opts createOptions

	cmd := &cobra.Command{
		Use:   "list [OPTIONS] NAMESPACE",
		Short: "List all document under a repo",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(args[0], &opts)
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

func runList(namespace string, opts *createOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	token := os.Getenv("token")
	c := yuque.NewClient(token)
	documents, err := c.Document.ListAll(namespace)
	if err != nil {
		return err
	}

	for _, d := range documents.Data {
		fmt.Printf("title: %s\n", d.Title)
	}

	return nil
}
