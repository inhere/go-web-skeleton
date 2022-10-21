package commands

import (
	"github.com/gookit/gcli/v3"
)

// InstallGoLintCommand Install lint/golint tool package
func InstallGoLintCommand() *gcli.Command {
	cmd := gcli.Command{
		Name:    "install:golint",
		Aliases: []string{"in:golint"},
		UseFor:  "Install go lint/golint tool package to local",
		Func: func(c *gcli.Command, args []string) error {
			return installToolPackage(toolPkgMap["golint"])
		},
		Help: `  How to use golint:
	golint ./...
`,
	}

	configInstallOption(&cmd)

	return &cmd
}
