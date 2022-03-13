package group

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
		Use:              "get GROUPLOGIN",
		Short:            "Get a group info",
		Args:             command.ExactArgs(1),
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(client, args[0])
		},
	}

	return cmd
}

func runGet(client *internal.Client, groupLogin string) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}
	groups, err := c.Group.GetInfo(groupLogin)
	if err != nil {
		return err
	}

	var t = `
type:               {{.Type}}
login:              {{.Login}}
name:               {{.Name}}
description:        {{.Description}}
avatarURL:          {{.AvatarUrl}}
reposCount:         {{.RepoCount}}
publicReposCount:   {{.PublicReposCount}}
publicTopicsCount:  {{.PublicTopicsCount}}
membersCount: 	    {{.MembersCount}}
grainsSum:          {{.GrainsSum}}
public:             {{.Public}}
followersCount:     {{.FollowersCount}},
followingCount":    {{.FollowingCount}}
createdAt:          {{.CreatedAt}},
updatedAt:          {{.UpdatedAt}},
	`
	groupInfoTemplate, err := template.New("t").Parse(t)
	if err != nil {
		return err
	}
	groupInfo := struct {
		Type              string
		Login             string
		Name              string
		Description       string
		Public            int
		AvatarUrl         string
		RepoCount         int
		PublicReposCount  int
		PublicTopicsCount int
		MembersCount      int
		GrainsSum         int
		FollowersCount    int
		FollowingCount    int
		CreatedAt         string
		UpdatedAt         string
	}{
		Type:              groups.Data.Type,
		Login:             groups.Data.Login,
		Name:              groups.Data.Name,
		Description:       groups.Data.Description,
		AvatarUrl:         groups.Data.AvatarURL,
		RepoCount:         groups.Data.BooksCount,
		PublicReposCount:  groups.Data.PublicBooksCount,
		PublicTopicsCount: groups.Data.PublicTopicsCount,
		MembersCount:      groups.Data.MembersCount,
		GrainsSum:         groups.Data.GrainsSum,
		Public:            groups.Data.Public,
		FollowersCount:    groups.Data.FollowersCount,
		FollowingCount:    groups.Data.FollowingCount,
		CreatedAt:         groups.Data.CreatedAt,
		UpdatedAt:         groups.Data.UpdatedAt,
	}

	return groupInfoTemplate.Execute(os.Stdout, groupInfo)
}
