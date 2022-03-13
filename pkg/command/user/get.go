package user

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
		Use:              "get",
		Short:            "Get user info",
		Args:             command.NoArgs,
		PersistentPreRun: client.CheckLogin(),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGet(client)
		},
	}

	return cmd
}

func runGet(client *internal.Client) error {
	c, err := yuque.NewClient(client.Token)
	if err != nil {
		return err
	}

	user, err := c.User.GetInfo()
	if err != nil {
		return err
	}

	var (
		t = `
login:              {{.Login}}
name:        	    {{.Name}}
description: 	    {{.Description}}
followersCount:     {{.FollowersCount}}
followingCount:     {{.FollowingCount}}
avatarURL:          {{.AvatarURL}}
reposCount:         {{.ReposCount}}
publicReposCount:   {{.PublicReposCount}}
publicTopicsCount:  {{.PublicTopicsCount}}
public:             {{.Public}}
createdAt:   	    {{.CreatedAt}}
updatedAt:    	    {{.UpdatedAt}}
		`

		userInfo = struct {
			Login             string
			Name              string
			Description       string
			Public            int
			CreatedAt         string
			UpdatedAt         string
			AvatarURL         string
			ReposCount        int
			PublicReposCount  int
			PublicTopicsCount int
			FollowersCount    int
			FollowingCount    int
		}{
			Login:             user.Data.Login,
			Name:              user.Data.Name,
			Description:       user.Data.Description,
			Public:            user.Data.Public,
			CreatedAt:         user.Data.CreatedAt,
			UpdatedAt:         user.Data.UpdatedAt,
			AvatarURL:         user.Data.AvatarURL,
			ReposCount:        user.Data.BooksCount,
			PublicReposCount:  user.Data.PublicBooksCount,
			PublicTopicsCount: user.Data.PublicTopicsCount,
			FollowersCount:    user.Data.FollowersCount,
			FollowingCount:    user.Data.FollowingCount,
		}
	)

	userInfoTemplate, err := template.New("t").Parse(t)
	if err != nil {
		return err
	}

	return userInfoTemplate.Execute(os.Stdout, userInfo)
}
