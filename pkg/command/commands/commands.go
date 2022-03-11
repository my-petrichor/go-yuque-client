package commands

import (
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/my-Sakura/go-yuque-client/pkg/command/document"
	"github.com/my-Sakura/go-yuque-client/pkg/command/group"
	"github.com/my-Sakura/go-yuque-client/pkg/command/registry"
	"github.com/my-Sakura/go-yuque-client/pkg/command/repo"
	"github.com/my-Sakura/go-yuque-client/pkg/command/searcher"
	"github.com/my-Sakura/go-yuque-client/pkg/command/user"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command, yuqueCli command.Cli) {
	cmd.AddCommand(
		user.NewUserCommand(yuqueCli),

		document.NewDocumentCommand(yuqueCli),

		group.NewGroupCommand(yuqueCli),

		repo.NewRepoCommand(yuqueCli),

		searcher.NewSearcherCommand(yuqueCli),

		registry.NewLoginCommand(yuqueCli),
		registry.NewLogoutCommand(yuqueCli),
	)
}
