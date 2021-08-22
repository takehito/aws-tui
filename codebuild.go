package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
)

func getAllProjects() ([]types.Project, error) {
	cli := codebuild.NewFromConfig(cfg)
	resp, err := cli.ListProjects(context.TODO(), &codebuild.ListProjectsInput{})
	if err != nil {
		return nil, err
	}
	bgp, err := cli.BatchGetProjects(context.TODO(), &codebuild.BatchGetProjectsInput{
		Names: resp.Projects,
	})
	if err != nil {
		return nil, err
	}

	return bgp.Projects, nil
}
