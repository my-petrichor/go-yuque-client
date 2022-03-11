package commands

import (
	"github.com/my-Sakura/go-yuque-client/pkg/command/document"
	"github.com/my-Sakura/go-yuque-client/pkg/command/group"
	"github.com/my-Sakura/go-yuque-client/pkg/command/registry"
	"github.com/my-Sakura/go-yuque-client/pkg/command/repo"
	"github.com/my-Sakura/go-yuque-client/pkg/command/searcher"
	"github.com/my-Sakura/go-yuque-client/pkg/command/user"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		user.NewUserCommand(),

		document.NewDocumentCommand(),

		group.NewGroupCommand(),

		repo.NewRepoCommand(),

		searcher.NewSearcherCommand(),

		registry.NewLoginCommand(),
		registry.NewLogoutCommand(),
	)
}
