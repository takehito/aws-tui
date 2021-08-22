package main

import (
	"context"
	"flag"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app *tview.Application
	cfg aws.Config
)

func main() {
	profile := flag.String("profile", "", "aws profile")
	flag.Parse()

	var err error
	cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(*profile))
	if err != nil {
		panic(err)
	}

	projects, err := getAllProjects()
	if err != nil {
		panic(err)
	}

	codebuildView := createCodebuildprojectsView(projects)
	app = tview.NewApplication()
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlD:
			app.Stop()
		}
		return event
	})
	app.SetRoot(codebuildView, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
