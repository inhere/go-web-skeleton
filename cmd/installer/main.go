package main

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/gookit/gcli/v2"
)

// go get https://github.com/inhere/go-web-skeleton/cmd/installer
var installOpts = struct {
	fontName   string
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

	cmd.BoolOpt(&installOpts.visualMode, "visual", "v", false, "Prints the font name.")
	cmd.StrOpt(&installOpts.fontName, "font", "", "", "Choose a font name. Default is a random font.")
	cmd.BoolOpt(&installOpts.list, "list", "", false, "Lists all available fonts.")
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

func install(_ *gcli.Command, args []string)  error {
	gcli.Print("hello, in the alone command\n")

	color.Info.Print("- clone the inhere/go-web-skeleton from github")

	// fmt.Printf("%+v\n", cmd.Flags)
	fmt.Printf("opts %+v\n", installOpts)
	fmt.Printf("args is %v\n", args)

	return nil
}
