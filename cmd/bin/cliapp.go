package main

import (
	"github.com/gookit/cliapp"
	"github.com/gookit/cliapp/builtin"
	"github.com/inhere/go-webx/cmd"
	"runtime"
)

// for test run: go build ./demo/cliapp.go && ./cliapp
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cliapp.NewApp()
	app.Version = "1.0.3"
	app.Description = "this is my cli application"

	app.SetVerbose(cliapp.VerbDebug)
	// app.DefaultCmd("exampl")

	app.Add(cmd.GitCommand())
	// app.Add(cmd.ColorCommand())
	app.Add(builtin.GenShAutoComplete())
	// fmt.Printf("%+v\n", cliapp.CommandNames())
	app.Run()
}
