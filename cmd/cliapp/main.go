package main

import (
	"github.com/gookit/gcli"
	"github.com/gookit/gcli/builtin"
	"github.com/inhere/go-web-skeleton/cmd"
	"runtime"
)

// for test run: go build ./demo/cliapp.go && ./cliapp
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := gcli.NewApp()
	app.Version = "1.0.3"
	app.Description = "this is my cli application"

	app.SetVerbose(gcli.VerbDebug)
	// app.DefaultCmd("exampl")

	app.Add(cmd.GitCommand())
	// app.Add(cmd.ColorCommand())
	app.Add(builtin.GenAutoCompleteScript())
	// fmt.Printf("%+v\n", cliapp.CommandNames())
	app.Run()
}
