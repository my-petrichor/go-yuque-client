package main

import (
	"fmt"
	"log"

	"github.com/my-Sakura/go-yuque-client/internal"
	"github.com/my-Sakura/go-yuque-client/pkg/command"
	"github.com/my-Sakura/go-yuque-client/pkg/command/commands"
	"github.com/my-Sakura/go-yuque-client/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	config.Init()
	client := internal.NewClient(viper.GetString("token"))

	if err := runYuque(client); err != nil {
		log.Fatalln(err)
	}
}

func runYuque(client *internal.Client) error {
	cmd := &cobra.Command{
		Use:           "yuque [OPTIONS] COMMAND [ARG...]",
		Short:         "A simple yuque application manage tool",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return command.ShowHelp()(cmd, args)
			}
			return fmt.Errorf("yuque: '%s' is not a yuque command.\nSee 'yuque --help'", args[0])
		},
		DisableFlagsInUseLine: true,
	}

	cmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", "Location of client config file (default $HOME/.config.yaml)")
	// cmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	// cmd.PersistentFlags().MarkShorthandDeprecated("help", "please use --help")
	// cmd.PersistentFlags().Lookup("help").Hidden = true

	cobra.AddTemplateFunc("add", func(a, b int) int { return a + b })
	cobra.AddTemplateFunc("hasSubCommands", hasSubCommands)
	cobra.AddTemplateFunc("hasManagementSubCommands", hasManagementSubCommands)
	cobra.AddTemplateFunc("operationSubCommands", operationSubCommands)
	cobra.AddTemplateFunc("managementSubCommands", managementSubCommands)

	cmd.SetUsageTemplate(usageTemplate)
	cmd.SetHelpTemplate(helpTemplate)

	commands.AddCommands(client, cmd)
	DisableFlagsInUseLine(cmd)

	return cmd.Execute()
}

func VisitAll(root *cobra.Command, fn func(*cobra.Command)) {
	for _, cmd := range root.Commands() {
		VisitAll(cmd, fn)
	}
	fn(root)
}

func DisableFlagsInUseLine(cmd *cobra.Command) {
	VisitAll(cmd, func(ccmd *cobra.Command) {
		// do not add a `[flags]` to the end of the usage line.
		ccmd.DisableFlagsInUseLine = true
	})
}

func hasSubCommands(cmd *cobra.Command) bool {
	return len(operationSubCommands(cmd)) > 0
}

func operationSubCommands(cmd *cobra.Command) []*cobra.Command {
	cmds := []*cobra.Command{}
	for _, sub := range cmd.Commands() {
		if sub.IsAvailableCommand() && !sub.HasSubCommands() {
			cmds = append(cmds, sub)
		}
	}
	return cmds
}

func hasManagementSubCommands(cmd *cobra.Command) bool {
	return len(managementSubCommands(cmd)) > 0
}

func managementSubCommands(cmd *cobra.Command) []*cobra.Command {
	cmds := []*cobra.Command{}
	for _, sub := range cmd.Commands() {
		if sub.IsAvailableCommand() && sub.HasSubCommands() {
			cmds = append(cmds, sub)
		}
	}
	return cmds
}

var (
	usageTemplate = `Usage:

{{- if not .HasSubCommands}}  {{.UseLine}}{{end}}
{{- if .HasSubCommands}}  {{ .CommandPath}}{{- if .HasAvailableFlags}} [OPTIONS]{{end}} COMMAND{{end}}

{{if ne .Long ""}}{{ .Long | trim }}{{ else }}{{ .Short | trim }}{{end}}
	
{{- if .HasAvailableFlags}}

Options:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

{{- end}}

{{- if hasManagementSubCommands . }}

Management Commands:
	
{{- range managementSubCommands . }}
  {{rpad .Name .NamePadding }} {{.Short}}
{{- end}}
	
{{- end}}

{{- if hasSubCommands .}}
	
Commands:
	
{{- range operationSubCommands . }}
  {{rpad .Name .NamePadding }} {{.Short}}
{{- end}}

{{- end}}
	
{{- if .HasSubCommands }}
	
Run '{{.CommandPath}} COMMAND --help' for more information on a command.
{{- end}}
	`

	helpTemplate = `
{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`
)
