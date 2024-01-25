package cli

import (
	"github.com/coder/coder/v2/cli/clibase"
	"github.com/coder/coder/v2/codersdk"
)

func (r *RootCmd) configure() *clibase.Cmd {
	client := new(codersdk.Client)

	cmd := &clibase.Cmd{
		Annotations: workspaceCommand,
		Use:         "configure",
		Short:       "Configure values required to connect to your Coder host.",
		Long: formatExamples(
			example{
				Description: "You can simply run this command to set the URL for your Coder host.",
				Command:     "coder configure",
			},
		),
		Middleware: clibase.Chain(clibase.RequireNArgs(0), r.InitClient(client)),
		Handler: func(inv *clibase.Invocation) error {
			out := inv.Stdout

			out.Write([]byte("arse"))

			return nil

		},
	}

	return cmd
}
