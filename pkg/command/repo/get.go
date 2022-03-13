package repo

import (
	"os"
	"text/template"

	yuque "github.com/my-Sakura/go-yuque-api"
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/spf13/cobra"
)

func newGetCommand(client *internal.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "get [OPTIONS] NAMESPACE",
		Short:            "Get repo info",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(client, args[0])
		},
	}

	return cmd
}

func runGet(client *internal.Client, namespace string) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	repo, err := c.Repo.GetInfo(namespace)
	if err != nil {
		return err
	}

	var (
		t = `
name:        	 {{.Name}}
description: 	 {{.Description}}
type:        	 {{.Kind}}
public:      	 {{.Public}}
createdAt:   	 {{.CreatedAt}}
updatedAt:    	 {{.UpdatedAt}}
itemsCount:   	 {{.ItemsCount}}
likesCount:   	 {{.LikesCount}}
watchedCount: 	 {{.WatchedCount}}
namespace:   	 {{.Namespace}}
slug:        	 {{.Slug}}
pinnedAt:     	 {{.PinnedAt}}
archived:     	 {{.Archived}}
userName:    	 {{.UserName}}        
userLogin:       {{.UserLogin}}
userDescription: {{.UserDescription}}
		`
		repoInfo = struct {
			Name            string
			Description     string
			Kind            string
			Public          int
			CreatedAt       string
			UpdatedAt       string
			ItemsCount      int
			LikesCount      int
			WatchedCount    int
			PinnedAt        string
			Archived        string
			Namespace       string
			Slug            string
			UserLogin       string
			UserName        string
			UserDescription string
		}{
			Name:            repo.Data.Name,
			Description:     repo.Data.Description,
			Namespace:       repo.Data.Namespace,
			Kind:            repo.Data.Type,
			Slug:            repo.Data.Slug,
			Public:          repo.Data.Public,
			ItemsCount:      repo.Data.ItemsCount,
			LikesCount:      repo.Data.LikesCount,
			WatchedCount:    repo.Data.WatchesCount,
			PinnedAt:        repo.Data.PinnedAt,
			Archived:        repo.Data.Archived,
			UpdatedAt:       repo.Data.UpdatedAt,
			CreatedAt:       repo.Data.CreatedAt,
			UserLogin:       repo.Data.User.Login,
			UserName:        repo.Data.User.Name,
			UserDescription: repo.Data.User.Description,
		}
	)

	repoInfoTemplate, err := template.New("t").Parse(t)
	if err != nil {
		return err
	}

	return repoInfoTemplate.Execute(os.Stdout, repoInfo)
}
