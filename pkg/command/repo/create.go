package repo

import (
	"errors"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type createOptions struct {
	slug        string
	kind        string
	name        string
	description string
	public      int
	userOrGroup int
}

func newCreateCommand(client *internal.Client) *cobra.Command {
	var opts createOptions

	cmd := &cobra.Command{
		Use:              "create [OPTIONS]",
		Short:            "Create a repo",
		Args:             command.NoArgs,
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(client, &opts)
		},
	}

	flags := cmd.Flags()
	flags.IntVar(&opts.userOrGroup, "user_or_group", 0, "Create repo under user or group (0 - user, 1 - group)")
	flags.StringVarP(&opts.slug, "slug", "s", "", "Slug of repo")
	flags.StringVarP(&opts.kind, "type", "t", "Book", "Type of repo (Book, Design, Sheet, Column, Resource, Thread)")
	flags.StringVarP(&opts.name, "name", "n", "", "Name of repo")
	flags.StringVarP(&opts.description, "description", "d", "", "Description of repo")
	flags.IntVarP(&opts.public, "public", "p", 0,
		"Public of repo (0 - private, 1 - all user, 2 - space member, 3 - all user under space (include external contact), 4 - only repo)")

	cmd.MarkFlagRequired("user_or_group")
	cmd.MarkFlagRequired("type")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("slug")

	return cmd
}

func runCreate(client *internal.Client, opts *createOptions) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	yuqueOption := yuque.RepoOption{
		Description: opts.description,
		Public:      opts.public,
	}
	if opts.userOrGroup == 0 {
		_, err = c.Repo.CreateUnderUser(opts.slug, opts.kind, opts.name, yuqueOption)
	} else if opts.userOrGroup == 1 {
		_, err = c.Repo.CreateUnderGroup(opts.slug, opts.kind, opts.name, yuqueOption)
	} else {
		return errors.New("Error flag userOrGroup")
	}

	return err
}
