package target

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/fullpipe/herald/place"
	"github.com/urfave/cli/v2"
)

type Cli struct {
	colors  bool
	message string
}

func (c *Cli) Name() string {
	return "cli"
}

func (c *Cli) Usage() string {
	return "cli"
}

func (c *Cli) Send(m place.Metadata) error {
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	if c.colors {
		if m.Project != "" {
			m.Project = yellow(m.Project)
		}

		if m.URL != "" {
			m.URL = cyan(m.URL)
		}

		if m.Branch != "" {
			m.Branch = cyan(m.Branch)
		}

		if m.SHA != "" {
			m.SHA = cyan(m.SHA)
		}

		if m.Author != "" {
			m.Author = cyan(m.Author)
		}

		if m.URL != "" {
			m.URL = cyan(m.URL)
		}
	}

	message, err := RenderMessage(c.message, m)
	if err != nil {
		return err
	}

	fmt.Print(message)

	return nil
}

func (c *Cli) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "message",
			Aliases: []string{"m"},
			Usage:   "Cli message template. Use metadata vars: {{.project}}, ...",
			Value: `
Project {{.project}} deployed

Project: {{ or .project "none"}}
URL: {{ or .url "none"}}
Branch: {{ or .branch "none"}}
CommitSHA: {{ or .sha "none"}}
Author: {{ or .author "none"}}
`,
			Destination: &c.message,
		},
		&cli.BoolFlag{
			Name:        "color",
			Aliases:     []string{"c"},
			Usage:       "Want some color?",
			Destination: &c.colors,
		},
	}
}
