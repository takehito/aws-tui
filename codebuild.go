package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
)

func getBuildsDescription(ids []string) ([]types.Build, error) {
	cli := codebuild.NewFromConfig(cfg)
	bgb, err := cli.BatchGetBuilds(context.Background(), &codebuild.BatchGetBuildsInput{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	return bgb.Builds, nil
}
