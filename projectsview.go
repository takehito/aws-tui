package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getBuilds(table *tview.TextView, projects []types.Project) func(index int, mainText string, secondaryText string, shortcut rune) {
	return func(index int, mainText string, secondaryText string, shortcut rune) {
		table.SetText(fmtProjectDescription(projects[index]))
	}
}

func createCodebuildprojectsView(projects []types.Project) tview.Primitive {
	table := tview.NewTextView()
	table.SetBorder(true)
	table.SetText(fmtProjectDescription(projects[0]))
	codebuilds := tview.NewList().
		ShowSecondaryText(false).
		SetDoneFunc(func() {
			table.Clear()
			app.SetFocus(table)
		}).SetSelectedBackgroundColor(tcell.Color102)
	for _, v := range projects {
		var desc string
		if v.Description != nil {
			desc = *v.Description
		}
		initial := []rune(strings.ToLower(*v.Name))[0]
		codebuilds.AddItem(*v.Name, desc, initial, func() {})
	}
	codebuilds.SetChangedFunc(getBuilds(table, projects))
	codebuilds.SetSelectedFunc(func(i int, _ string, _ string, _ rune) {
		cli := codebuild.NewFromConfig(cfg)
		builds, err := cli.ListBuildsForProject(context.Background(), &codebuild.ListBuildsForProjectInput{
			ProjectName: projects[i].Name,
		})
		if err != nil {
			panic(err)
		}
		if len(builds.Ids) == 0 {
			return
		}

		app.SetRoot(createBuildsView(builds.Ids), true)
	})
	codebuilds.SetTitle("codebuild projects").SetBorder(true)

	return tview.NewFlex().
		AddItem(codebuilds, 0, 1, true).
		AddItem(table, 0, 2, true)
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
