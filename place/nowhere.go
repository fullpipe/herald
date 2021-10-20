package place

import "github.com/urfave/cli/v2"

type Nowhere struct {
	project string
	url     string
	branch  string
	sha     string
	author  string
}

func (t *Nowhere) Name() string {
	return "nowhere"
}

func (t *Nowhere) Usage() string {
	return ""
}

func (t *Nowhere) Metadata() (Metadata, error) {
	m := Metadata{
		Project: t.project,
		URL:     t.url,
		Branch:  t.branch,
		SHA:     t.sha,
		Author:  t.author,
	}

	return m, nil
}

func (t *Nowhere) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "project",
			Usage:       "",
			Destination: &t.project,
		},
		&cli.StringFlag{
			Name:        "url",
			Usage:       "",
			Destination: &t.url,
		},

		&cli.StringFlag{
			Name:        "branch",
			Usage:       "",
			Destination: &t.branch,
		},
		&cli.StringFlag{
			Name:        "sha",
			Usage:       "",
			Destination: &t.sha,
		},
		&cli.StringFlag{
			Name:        "author",
			Usage:       "",
			Destination: &t.author,
		},
	}
}
