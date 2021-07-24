package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application

func main() {
	profile := flag.String("profile", "", "aws profile")
	flag.Parse()

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(*profile))
	if err != nil {
		panic(err)
	}

	cli := codebuild.NewFromConfig(cfg)
	resp, err := cli.ListProjects(context.TODO(), &codebuild.ListProjectsInput{})
	bgp, err := cli.BatchGetProjects(context.TODO(), &codebuild.BatchGetProjectsInput{
		Names: resp.Projects,
	})
	if err != nil {
		panic(err)
	}

	projects := bgp.Projects
	table := tview.NewTextView()
	table.SetBorder(true)
	table.SetText(fmtProjectDescription(projects[0]))
	codebuilds := tview.NewList().
		ShowSecondaryText(false).
		SetDoneFunc(func() {
			table.Clear()
			app.SetFocus(table)
		}).SetSelectedBackgroundColor(tcell.Color102)
	for _, v := range bgp.Projects {
		var desc string
		if v.Description != nil {
			desc = *v.Description
		}
		codebuilds.AddItem(*v.Name, desc, []rune(*v.Name)[0], func() {})
	}
	codebuilds.SetChangedFunc(getBuilds(table, bgp.Projects))
	codebuilds.SetTitle("codebuild projects").SetBorder(true)

	flex := tview.NewFlex().
		AddItem(codebuilds, 0, 1, true).
		AddItem(table, 0, 2, true)

	app = tview.NewApplication()
	app.SetRoot(flex, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func getBuilds(table *tview.TextView, projects []types.Project) func(index int, mainText string, secondaryText string, shortcut rune) {
	return func(index int, mainText string, secondaryText string, shortcut rune) {
		table.SetText(fmtProjectDescription(projects[index]))
	}
}

func fmtProjectDescription(project types.Project) string {
	text := strings.Builder{}
	if project.Name != nil {
		txt := fmt.Sprintf("name: %s\n", *project.Name)
		text.WriteString(txt)
	}
	if project.Arn != nil {
		txt := fmt.Sprintf("arn: %s\n", *project.Arn)
		text.WriteString(txt)
	}
	text.WriteString(fmt.Sprintf("created: %s\n", project.Created))
	text.WriteString(fmt.Sprintf("last modified: %s\n", project.LastModified))
	if project.ServiceRole != nil {
		txt := fmt.Sprintf("service role: %s\n", *project.ServiceRole)
		text.WriteString(txt)
	}
	if project.ConcurrentBuildLimit != nil {
		txt := fmt.Sprintf("conccurent build limit: %d\n", *project.ConcurrentBuildLimit)
		text.WriteString(txt)
	}
	if project.Description != nil {
		txt := fmt.Sprintf("description: %s\n", *project.Description)
		text.WriteString(txt)
	}
	if project.EncryptionKey != nil {

		txt := fmt.Sprintf("encryption key: %s\n", *project.EncryptionKey)
		text.WriteString(txt)
	}
	if project.QueuedTimeoutInMinutes != nil {
		txt := fmt.Sprintf("queued timeout in minutes: %d\n", *project.QueuedTimeoutInMinutes)
		text.WriteString(txt)
	}
	if project.SourceVersion != nil {
		txt := fmt.Sprintf("source version: %s\n", *project.SourceVersion)
		text.WriteString(txt)
	}
	text.WriteString(fmt.Sprintf("build spec: \n%s\n", *project.Source.Buildspec))

	return text.String()
}
