package place

import (
	"github.com/urfave/cli/v2"
)

type Gitlab struct {
	project string
	url     string
	branch  string
	sha     string
	author  string
}

func (g *Gitlab) Name() string {
	return "gitlab"
}

func (g *Gitlab) Usage() string {
	return "collect metadata from gitlab ci envars"
}

func (g *Gitlab) Metadata() (Metadata, error) {
	m := Metadata{
		Project: g.project,
		URL:     g.url,
		Branch:  g.branch,
		SHA:     g.sha,
		Author:  g.author,
	}

	//PipelineURL???
	// m.URL = os.Getenv("CI_PIPELINE_URL")

	return m, nil
}

func (g *Gitlab) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "project",
			EnvVars:     []string{"CI_PROJECT_TITLE"},
			Usage:       "",
			Destination: &g.project,
		},
		&cli.StringFlag{
			Name:        "url",
			EnvVars:     []string{"CI_PROJECT_URL"},
			Usage:       "",
			Destination: &g.url,
		},

		&cli.StringFlag{
			Name:        "branch",
			EnvVars:     []string{"CI_COMMIT_BRANCH"},
			Usage:       "",
			Destination: &g.branch,
		},
		&cli.StringFlag{
			Name:        "sha",
			EnvVars:     []string{"CI_COMMIT_SHA"},
			Usage:       "",
			Destination: &g.sha,
		},
		&cli.StringFlag{
			Name:        "author",
			EnvVars:     []string{"CI_COMMIT_AUTHOR"},
			Usage:       "",
			Destination: &g.author,
		},
	}
}
