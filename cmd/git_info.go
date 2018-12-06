package cmd

import (
	"fmt"
	cli "github.com/gookit/cliapp"
	cliutil "github.com/gookit/cliapp/utils"
	"github.com/gookit/color"
	"github.com/inhere/go-wex-skeleton/app/utils"
	"github.com/inhere/go-wex-skeleton/model"
	"log"
	"strings"
)

var gitOpts GitOpts

type GitOpts struct {
	output string
}

// GitCommand
func GitCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "git",
		Aliases: []string{"git-info"},
		UseFor:  "collect project info by git info",
		Func: gitExecute,
	}

	gitOpts = GitOpts{}

	cmd.StrOpt(&gitOpts.output, "output", "o", "static/app.json", "output file of the git info")

	return &cmd
}

// arg test:
// 	go build cliapp.go && ./cliapp git --id 12 -c val ag0 ag1
func gitExecute(cmd *cli.Command, args []string) int {
	info := model.GitInfoData{}

	// latest commit id by: git log --pretty=%H -n1 HEAD
	cid, err := cliutil.ExecCommand("git log --pretty=%H -n1 HEAD")
	if err != nil {
		log.Fatal(err)
		return -2
	}

	cid = strings.TrimSpace(cid)
	fmt.Printf("commit id: %s\n", cid)
	info.Version = cid

	// latest commit date by: git log -n1 --pretty=%ci HEAD
	cDate, err := cliutil.ExecCommand("git log -n1 --pretty=%ci HEAD")
	if err != nil {
		log.Fatal(err)
		return -2
	}

	cDate = strings.TrimSpace(cDate)
	info.ReleaseAt = cDate
	fmt.Printf("commit date: %s\n", cDate)

	// get tag: git describe --tags --exact-match HEAD
	tag, err := cliutil.ExecCommand("git describe --tags --exact-match HEAD")
	if err != nil {
		// get branch: git branch -a | grep "*"
		br, err := cliutil.ExecCommand(`git branch -a | grep "*"`)
		if err != nil {
			log.Fatal(err)
			return -2
		}
		br = strings.TrimSpace(strings.Trim(br, "*"))
		info.Tag = br
		fmt.Printf("current branch: %s\n", br)
	} else {
		tag = strings.TrimSpace(tag)
		info.Tag = tag
		fmt.Printf("latest tag: %s\n", tag)
	}

	err = utils.WriteJsonFile(gitOpts.output, &info)

	if err != nil {
		log.Fatal(err)
		return -2
	}

	color.FgGreen.Println("\nOk, project info collect completed!")
	return 0
}
