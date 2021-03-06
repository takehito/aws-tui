package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

func getLogGroups(logGroupPrefix string) ([]types.LogGroup, error) {
	cli := cloudwatchlogs.NewFromConfig(cfg)
	dlg, err := cli.DescribeLogGroups(context.Background(), &cloudwatchlogs.DescribeLogGroupsInput{
		LogGroupNamePrefix: &logGroupPrefix,
	})
	if err != nil {
		return nil, err
	}
	return dlg.LogGroups, nil
}

func getLog(groupName string, streamName string) ([]types.OutputLogEvent, error) {
	cli := cloudwatchlogs.NewFromConfig(cfg)
	gle, err := cli.GetLogEvents(context.Background(), &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  &groupName,
		LogStreamName: &streamName,
	})
	if err != nil {
		return nil, err
	}
	return gle.Events, nil
}
