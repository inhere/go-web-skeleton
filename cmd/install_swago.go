package cmd

import (
	"github.com/gookit/gcli/v2"
	"github.com/gookit/goutil/sysutil"
)

var toolPkgMap = map[string]string{
	"swag": "github.com/swaggo/swag/cmd/swag",
	// lint tools
	"revive": "github.com/mgechev/revive",
	"golint": "golang.org/x/lint/golint",
}

// InstallSwagCommand Install swaggo/swag tool package
func InstallSwagCommand() *gcli.Command {
	cmd := gcli.Command{
		Name:    "install:swag",
		Aliases: []string{"in:swag"},
		UseFor:  "collect project info by git info",
		Func: func(c *gcli.Command, args []string) error {
			return installToolPackage(toolPkgMap["swag"])
		},
	}

	cmd.StrOpt(
		&gitOpts.output, "output", "o", "static/app.json",
		"output file of the git info",
	)

	return &cmd
}

func installToolPackage(pkgUrl string) error {
	cmdArgs := []string{
		"-u",
		"-v",
		pkgUrl,
	}

	// eg: go get -u github.com/swaggo/swag/cmd/swag
	_, err := sysutil.ExecCmd("go", cmdArgs, "/tmp")
	return err
}
