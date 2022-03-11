package repo

import (
	"os"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type updateOptions struct {
	namespace   string
	slug        string
	kind        string
	name        string
	description string
	public      int
}

func newUpdateCommand() *cobra.Command {
	var opts updateOptions

	cmd := &cobra.Command{
		Use:   "update [OPTIONS]",
		Short: "Update a repo (must set namespace flag)",
		Args:  command.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate(&opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.namespace, "namespace", "n", "", "Namespace of repo")
	flags.StringVarP(&opts.slug, "slug", "s", "", "Slug of repo")
	flags.StringVarP(&opts.kind, "type", "t", "Book", "Type of repo (Book, Design, Sheet, Column, Resource, Thread) default Book")
	flags.StringVar(&opts.name, "name", "", "Name of repo")
	flags.StringVarP(&opts.description, "description", "d", "", "Description of repo")
	flags.IntVarP(&opts.public, "public", "p", 0,
		"Public of repo (0 - private, 1 - all user, 2 - space member, 3 - all user under space (include external contact), 4 - only repo) default 0")

	return cmd
}

func runUpdate(opts *updateOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }

	token := os.Getenv("token")
	c := yuque.NewClient(token)
	yuqueOption := yuque.RepoOption{
		Slug:        opts.slug,
		Name:        opts.name,
		Description: opts.description,
		Public:      opts.public,
		Kind:        opts.kind,
	}
	_, err := c.Repo.Update(opts.namespace, yuqueOption)

	return err
}
