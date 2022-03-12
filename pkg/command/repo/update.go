package repo

import (
	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type updateOptions struct {
	slug        string
	kind        string
	name        string
	description string
	public      int
}

func newUpdateCommand(client *internal.Client) *cobra.Command {
	var opts updateOptions

	cmd := &cobra.Command{
		Use:   "update [OPTIONS] NAMESPACE",
		Short: "Update a repo",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate(client, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.slug, "slug", "s", "", "Slug of repo")
	flags.StringVarP(&opts.kind, "type", "t", "Book", "Type of repo (Book, Design, Sheet, Column, Resource, Thread) default Book")
	flags.StringVar(&opts.name, "name", "", "Name of repo")
	flags.StringVarP(&opts.description, "description", "d", "", "Description of repo")
	flags.IntVarP(&opts.public, "public", "p", 0,
		"Public of repo (0 - private, 1 - all user, 2 - space member, 3 - all user under space (include external contact), 4 - only repo) default 0")

	return cmd
}

func runUpdate(client *internal.Client, namespace string, opts *updateOptions) error {
	if !client.IsLogin() {
		return internal.ErrNoLogin
	}

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	yuqueOption := yuque.RepoOption{
		Slug:        opts.slug,
		Name:        opts.name,
		Description: opts.description,
		Public:      opts.public,
		Kind:        opts.kind,
	}
	_, err = c.Repo.Update(namespace, yuqueOption)

	return err
}
