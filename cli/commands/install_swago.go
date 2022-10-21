package commands

import (
	"github.com/gookit/gcli/v3"
)

// InstallSwagCommand Install swaggo/swag tool package
func InstallSwagCommand() *gcli.Command {
	cmd := gcli.Command{
		Name:    "install:swag",
		Aliases: []string{"in:swag"},
		UseFor:  "Install swaggo/swag tool package to local",
		Func: func(c *gcli.Command, args []string) error {
			return installToolPackage(toolPkgMap["swag"])
		},
		Help: `  How to use swag:
	swag init -s static
`,
	}

	configInstallOption(&cmd)

	return &cmd
}
