package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/rivo/tview"
)

func createLogView(log []types.OutputLogEvent) tview.Primitive {
	var message strings.Builder
	for _, v := range log {
		line := fmt.Sprintf("%s\n", *v.Message)
		_, err := message.WriteString(line)
		if err != nil {
			panic(err)
		}
	}
	return tview.NewTextView().SetText(message.String())
}
