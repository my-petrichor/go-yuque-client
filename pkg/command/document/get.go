package document

import (
	"os"
	"text/template"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

type getOptions struct {
}

func newGetCommand(client *internal.Client) *cobra.Command {
	var opts getOptions

	cmd := &cobra.Command{
		Use:              "get [OPTIONS] PATH",
		Short:            "Get user info",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(client, args[0], &opts)
		},
	}

	return cmd
}

func runGet(client *internal.Client, path string, opts *getOptions) error {
	namespace, slug, err := splitPath(path)
	if err != nil {
		return err
	}

	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	doc, err := c.Document.GetInfo(namespace, slug)
	if err != nil {
		return err
	}

	var t = `
format:          {{.Format}}			 
slug:            {{.Slug}}
title:           {{.Title}}
wordCount:       {{.WordCount}}
public:          {{.Public}}
likeCount:       {{.LikeCount}}
commentsCount:	 {{.CommentsCount}}
createdAt:       {{.CreatedAt}}
firstPublished:	 {{.FirstPublishedAt}}
contentUpdated:	 {{.ContentUpdatedAt}}
updatedAt:       {{.UpdatedAt}}
publishedAt:	 {{.PublishedAt}}
repoName:        {{.RepoName}}
repoNamespace:	 {{.RepoNamespace}}
userName:        {{.UserName}}
userLogin:       {{.UserLogin}}
creatorName:	 {{.CreatorName}}
creatorLogin:	 {{.CreatorLogin}}
		`

	docInfoTemplate, err := template.New("t").Parse(t)
	if err != nil {
		return err
	}

	docInfo := struct {
		Format string
		Slug   string
		Title  string
		// Content          string
		WordCount        int
		Public           int
		LikeCount        int
		CommentsCount    int
		CreatedAt        string
		FirstPublishedAt string
		ContentUpdatedAt string
		UpdatedAt        string
		PublishedAt      string
		RepoName         string
		RepoNamespace    string
		UserName         string
		UserLogin        string
		CreatorName      string
		CreatorLogin     string
	}{
		Format: doc.Data.Format,
		Slug:   doc.Data.Slug,
		Title:  doc.Data.Title,
		// Content:          doc.Data.Body,
		WordCount:        doc.Data.WordCount,
		Public:           doc.Data.Public,
		LikeCount:        doc.Data.LikeCount,
		CommentsCount:    doc.Data.CommentsCount,
		CreatedAt:        doc.Data.CreatedAt,
		FirstPublishedAt: doc.Data.FirstPublishedAt,
		ContentUpdatedAt: doc.Data.ContentUpdatedAt,
		UpdatedAt:        doc.Data.UpdatedAt,
		PublishedAt:      doc.Data.PublishedAt,
		RepoName:         doc.Data.Book.Name,
		RepoNamespace:    doc.Data.Book.Namespace,
		UserName:         doc.Data.Book.User.Name,
		UserLogin:        doc.Data.Book.User.Login,
		CreatorName:      doc.Data.Creator.Name,
		CreatorLogin:     doc.Data.Creator.Login,
	}

	return docInfoTemplate.Execute(os.Stdout, docInfo)
}
