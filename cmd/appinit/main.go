package main

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/goutil/sysutil"
)

// creator
// go get https://github.com/inhere/go-web-skeleton/cmd/creator
// installer
// go get https://github.com/inhere/go-web-skeleton/cmd/installer
var installOpts = struct {
	dirName    string
	fullName   string
	visualMode bool
	list       bool
	sample     bool
}{}

// test run: go build ./_examples/alone && ./alone -h
func main() {
	cmd := gcli.Command{
		Name:    "install",
		Aliases: []string{"ts"},
		UseFor:  "this is a description <info>message</> for {$cmd}", // // {$cmd} will be replace to 'test'
	}

	cmd.BoolOpt(
		&installOpts.visualMode,
		"visual", "v", false,
		"Prints the font name.",
	)
	cmd.StrOpt(
		&installOpts.dirName,
		"name", "n", "",
		"Choose a font name. Default is a random font.",
	)
	cmd.StrOpt(
		&installOpts.fullName,
		"full-name", "", "",
		"Choose a font name. Default is a random font.",
	)
	cmd.BoolOpt(
		&installOpts.list,
		"list", "", false,
		"Lists all available fonts.",
	)
	cmd.BoolOpt(
		&installOpts.sample,
		"sample",
		"",
		false,
		"Prints a sample with that font.",
	)

	cmd.Func = install

	// Alone Running
	cmd.MustRun(nil)
}

func install(_ *gcli.Command, args []string) error {
	gcli.Print("hello, in the alone command\n")

	color.Warn.Println("- clone the inhere/go-web-skeleton from github")
	_, err := sysutil.QuickExec("git clone https://github.com/inhere/go-web-skeleton")
	if err != nil {
		return err
	}

	// replace "github.com/inhere/go-web-skeleton" to your package name
	// dirs := []string{"api", "app", "model"}

	// fmt.Printf("%+v\n", cmd.Flags)
	fmt.Printf("opts %+v\n", installOpts)
	fmt.Printf("args is %v\n", args)

	return nil
}
