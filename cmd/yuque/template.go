package main

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
