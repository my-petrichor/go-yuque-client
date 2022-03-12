package searcher

import (
	"fmt"

	huge "github.com/dablelv/go-huge-util"
	"github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type searcherOptions struct {
	kind   string
	offset int
}

func NewSearcherCommand(client *internal.Client) *cobra.Command {
	var opts searcherOptions

	cmd := &cobra.Command{
		Use:              "searcher [OPTIONS] KEYWORD",
		Short:            "Search by keyword",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSearcher(client, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.kind, "type", "t", "doc", "Type of search (doc, repo, artboard, group, user, attachment), default doc")
	flags.IntVarP(&opts.offset, "offset", "o", 1, "Offset of search, default 1")

	return cmd
}

func runSearcher(client *internal.Client, keyWord string, opts *searcherOptions) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	s, err := c.Search.Work(keyWord, yuque.SearcherOption{
		Kind:   opts.kind,
		Offset: opts.offset,
	})
	if err != nil {
		return err
	}

	data, err := huge.ToIndentJSON(&s.Data)
	if err != nil {
		return err
	}
	fmt.Println(data)

	return nil
}
