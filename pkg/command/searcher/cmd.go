package searcher

import (
	"fmt"
	"os"

	"github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type searcherOptions struct {
	kind string
}

func NewSearcherCommand(yuqueCli command.Cli) *cobra.Command {
	var opts searcherOptions

	cmd := &cobra.Command{
		Use:   "searcher",
		Short: "Search by keyword",
		Args:  command.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSearcher(yuqueCli, args[0], &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.kind, "type", "t", "", "search type")

	return cmd
}

func runSearcher(yuqueCli command.Cli, keyWord string, opts *searcherOptions) error {
	// if !command.Login() {
	// 	return internal.ErrNoLogin
	// }
	var (
		kind  = opts.kind
		token = os.Getenv("token")
	)

	if opts.kind == "" {
		kind = "doc"
	}

	c := yuque.NewClient(token)
	s, err := c.Search.Work(kind, keyWord)
	if err != nil {
		return err
	}
	for _, v := range s.Data {
		fmt.Println(v.Title, v.Summary)
	}

	return nil
}
