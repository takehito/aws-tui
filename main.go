package main

import (
	"context"
	"flag"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/rivo/tview"
)

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

	list := tview.NewList()
	for _, v := range bgp.Projects {
		var desc string
		if v.Description != nil {
			desc = *v.Description
		}
		list.AddItem(*v.Name, desc, []rune(*v.Name)[0], func() {})
	}

	list.SetTitle("codebuild projects").SetBorder(true)

	app := tview.NewApplication().SetRoot(list, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
