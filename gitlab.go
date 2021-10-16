package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

type Gitlab struct{}

func (g *Gitlab) Name() string {
	return "gitlab"
}

func (g *Gitlab) Metadata() (Metadata, error) {
	m := Metadata{}

	m.ProjectTitle = os.Getenv("CI_PROJECT_TITLE")
	m.ProjectURL = os.Getenv("CI_PROJECT_URL")

	m.Branch = os.Getenv("CI_COMMIT_BRANCH")
	m.CommitSHA = os.Getenv("CI_COMMIT_SHA")
	m.Author = os.Getenv("CI_COMMIT_AUTHOR")

	//PipelineURL???
	m.URL = os.Getenv("CI_PIPELINE_URL")

	return m, nil
}

func (g *Gitlab) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "title",
			Usage: "Want some color?",
		},
	}
}
