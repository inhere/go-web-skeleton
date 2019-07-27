package cmd

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/gcli/v2"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/jsonutil"
	"github.com/inhere/go-web-skeleton/model"
)

var gitOpts = &struct {
	output string
}{}

// GitCommand
func GitCommand() *gcli.Command {
	cmd := gcli.Command{
		Name:    "git",
		Aliases: []string{"git-info"},
		UseFor:  "collect project info by git info",
		Func:    gitExecute,
	}

	cmd.StrOpt(&gitOpts.output, "output", "o", "static/app.json", "output file of the git info")

	return &cmd
}

// arg test:
// 	go build cliapp.go && ./cliapp git --id 12 -c val ag0 ag1
func gitExecute(_ *gcli.Command, _ []string) (err error) {
	info := model.GitInfoData{}

	// latest commit id by: git log --pretty=%H -n1 HEAD
	cid, err := cliutil.ShellExec("git log --pretty=%H -n1 HEAD")
	if err != nil {
		return err
	}

	cid = strings.TrimSpace(cid)
	fmt.Printf("commit id: %s\n", cid)
	info.Version = cid

	// latest commit date by: git log -n1 --pretty=%ci HEAD
	cDate, err := cliutil.ShellExec("git log -n1 --pretty=%ci HEAD")
	if err != nil {
		return err
	}

	cDate = strings.TrimSpace(cDate)
	info.ReleaseAt = cDate
	fmt.Printf("commit date: %s\n", cDate)

	// get tag: git describe --tags --exact-match HEAD
	tag, err := cliutil.ShellExec("git describe --tags --exact-match HEAD")
	if err != nil {
		// get branch: git branch -a | grep "*"
		br, err := cliutil.ShellExec(`git branch -a | grep "*"`)
		if err != nil {
			return err
		}

		br = strings.TrimSpace(strings.Trim(br, "*"))
		info.Tag = br
		fmt.Printf("current branch: %s\n", br)
	} else {
		tag = strings.TrimSpace(tag)
		info.Tag = tag
		fmt.Printf("latest tag: %s\n", tag)
	}

	err = jsonutil.WriteFile(gitOpts.output, &info)
	if err != nil {
		return
	}

	color.FgGreen.Println("\nOk, project info collect completed!")
	return
}
