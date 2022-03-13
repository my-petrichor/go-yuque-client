package searcher

import (
	"os"
	"text/template"

	"github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type searcherOptions struct {
	kind   string
	offset int
	simple int
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
	flags.StringVarP(&opts.kind, "type", "t", "doc", "Type of search (doc, repo, artboard, group, user, attachment)")
	flags.IntVarP(&opts.offset, "offset", "o", 1, "Offset of search")
	flags.IntVarP(&opts.simple, "simple", "s", 0, "Simple output of search")

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

	var t string
	if opts.simple == 1 {
		t = `
{{.ID}}.	
info:  {{.Info}}
url:   {{.Url}}
		`
	} else {
		t = `
{{.ID}}.	
info:             {{.Info}}
url:              {{.Url}}
userLogin: 	      {{.UserLogin}}
userName:         {{.UserName}}
userDescription:  {{.UserDescription}}
repoSlug: 	      {{.RepoSlug}}
repoName: 		  {{.RepoName}}
repoDescription:  {{.RepoDescription}}
docSlug:          {{.DocSlug}}
docTitle:         {{.DocTitle}}
docDescription:   {{.DocDescription}}
`
	}

	searchInfoTemplate, err := template.New("t").Parse(t)
	if err != nil {
		return err
	}

	for i, v := range s.Data {
		searchInfo := struct {
			ID              int
			Info            string
			Url             string
			UserLogin       string
			UserName        string
			UserDescription interface{}
			RepoSlug        string
			RepoName        string
			RepoDescription string
			DocSlug         string
			DocTitle        string
			DocDescription  string
		}{
			ID:              i + 1,
			Info:            v.Info,
			Url:             v.URL,
			UserLogin:       v.Target.Book.User.Login,
			UserName:        v.Target.Book.User.Name,
			UserDescription: v.Target.Book.User.Description,
			RepoSlug:        v.Target.Book.Slug,
			RepoName:        v.Target.Book.Name,
			RepoDescription: v.Target.Book.Description,
			DocSlug:         v.Target.Slug,
			DocTitle:        v.Target.Title,
			DocDescription:  v.Target.Description,
		}
		if err = searchInfoTemplate.Execute(os.Stdout, searchInfo); err != nil {
			return err
		}
	}

	return nil
}
