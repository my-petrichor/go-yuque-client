package commands

import (
	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command/document"
	"github.com/my-Sakura/go-yuque-client/pkg/command/group"
	"github.com/my-Sakura/go-yuque-client/pkg/command/registry"
	"github.com/my-Sakura/go-yuque-client/pkg/command/repo"
	"github.com/my-Sakura/go-yuque-client/pkg/command/searcher"
	"github.com/my-Sakura/go-yuque-client/pkg/command/user"
	"github.com/spf13/cobra"
)

func AddCommands(client *internal.Client, cmd *cobra.Command) {
	cmd.AddCommand(
		user.NewUserCommand(client),

		document.NewDocumentCommand(client),

		group.NewGroupCommand(client),

		repo.NewRepoCommand(client),

		searcher.NewSearcherCommand(client),

		registry.NewLoginCommand(client),
		registry.NewLogoutCommand(client),
	)
}
