package commands

import (
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/progress"
	"github.com/gookit/goutil/sysutil"
)

var toolPkgMap = map[string]string{
	"swag": "github.com/swaggo/swag/cmd/swag",
	// lint tools
	"revive": "github.com/mgechev/revive",
	"golint": "golang.org/x/lint/golint",
}

var installOpt = struct {
	update  bool
	verbose bool
}{}

func configInstallOption(c *gcli.Command) {
	c.BoolOpt(
		&installOpt.update, "update", "u", false,
		"update package, will add '-u' on go get",
	)
	c.BoolOpt(
		&installOpt.verbose, "verbose", "v", false,
		"update package, will add '-v' on go get",
	)
}

func installToolPackage(pkgUrl string) error {
	cmdArgs := []string{"get"}

	if installOpt.update {
		cmdArgs = append(cmdArgs, "-u")
	}

	if installOpt.verbose {
		cmdArgs = append(cmdArgs, "-v")
	}

	cmdArgs = append(cmdArgs, pkgUrl)
	cmdStr := strings.Join(cmdArgs, " ")
	color.Yellow.Println("> go ", cmdStr)

	s := progress.RoundTripSpinner(
		progress.GetCharTheme(0),
		time.Duration(100)*time.Millisecond,
	)

	// s.Start("%s work handling ... ...")
	s.Start("[%s] Work Handling")

	// eg: go get -v -u github.com/swaggo/swag/cmd/swag
	msg, err := sysutil.ExecCmd("go", cmdArgs, "/tmp")

	// s.Stop("work handle complete")
	s.Stop()

	if msg != "" {
		color.Red.Println(msg)
	}

	if err == nil {
		color.Green.Println("Install Completed!")
	}
	return err
}
