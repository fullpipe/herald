package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	places := []Place{&Nowhere{}}
	targets := []Target{&Cli{}}

	app := cli.NewApp()

	for _, place := range places {
		subCommands := []*cli.Command{}

		for _, target := range targets {
			subCommands = append(subCommands, &cli.Command{
				Name:  target.Name(),
				Usage: "add a new template",
				Flags: target.Flags(),
				Action: func(c *cli.Context) error {
					meta, err := place.Metadata()
					if err != nil {
						return err
					}

					fmt.Println(meta)

					return target.Send(meta)
				},
			})
		}

		app.Commands = append(app.Commands, &cli.Command{
			Name:        place.Name(),
			Usage:       "options for task templates",
			Flags:       place.Flags(),
			Subcommands: subCommands,
		})
	}

	// app := &cli.App{
	// 	Name:  "herald",
	// 	Usage: "notify someone from somewhere",
	// 	Flags: []cli.Flag{
	// 		&cli.StringFlag{
	// 			Name:    "message",
	// 			Aliases: []string{"m"},
	// 			Usage:   "Message to send `MESSAGE`",
	// 		},
	// 	},
	// 	Action: func(c *cli.Context) error {
	// 		place, err := GetPlace(c.Args().First())
	// 		if err != nil {
	// 			return err
	// 		}

	// 		target, err := GetTarget(c.Args().Get(1))
	// 		if err != nil {
	// 			return err
	// 		}

	// 		meta, err := place.GetMetadata()
	// 		if err != nil {
	// 			return err
	// 		}

	// 		message, err := GetMessage(meta, c.String("message"))
	// 		if err != nil {
	// 			return err
	// 		}

	// 		return target.Send(message)
	// 	},
	// }

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPlace(p string) (Place, error) {
	switch p {
	case "gitlab":
		return &Gitlab{}, nil
	case "github":
		return nil, errors.New("No github implimentation")
	default:
		return nil, fmt.Errorf("No %s place implimentation", p)
	}
}

func GetMessage(m Metadata, template string) (Message, error) {
	return Message{}, nil
}

func GetTarget(t string) (Target, error) {
	switch t {
	case "slack":
		return &Slack{}, nil
	default:
		return nil, fmt.Errorf("No %s target implimentation", t)
	}
}

type Message struct {
	Text string
}

type Metadata struct {
	ProjectTitle string
	ProjectURL   string
	Branch       string
	CommitSHA    string
	Author       string
	URL          string
}

// ----
type Place interface {
	Name() string
	Metadata() (Metadata, error)
	Flags() []cli.Flag
}

type Nowhere struct {
	projectTitle string
	projectURL   string
	branch       string
	commitSHA    string
	author       string
	url          string
}

func (t *Nowhere) Name() string {
	return "nowhere"
}

func (t *Nowhere) Metadata() (Metadata, error) {
	m := Metadata{
		ProjectTitle: t.projectTitle,
		ProjectURL:   t.projectURL,
		Branch:       t.branch,
		CommitSHA:    t.commitSHA,
		Author:       t.author,
		URL:          t.url,
	}

	return m, nil
}

func (t *Nowhere) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "projectTitle",
			Usage:       "",
			Destination: &t.projectTitle,
		},
		&cli.StringFlag{
			Name:        "branch",
			Usage:       "",
			Destination: &t.branch,
		},
	}
}

// ----
type Target interface {
	Name() string
	Send(m Metadata) error
	Flags() []cli.Flag
}

type Cli struct {
	colors  bool
	message string
}

func (c *Cli) Name() string {
	return "cli"
}

func (c *Cli) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "message",
			Aliases: []string{"m"},
			Usage:   "Want some color?",
			Value: `
Project {{.ProjectTitle}} deployed
Project: {{ or .ProjectTitle "none"}}
URL: {{ or .ProjectURL "none"}}
Branch: {{ or .Branch "none"}}
CommitSHA: {{ or .CommitSHA "none"}}
Author: {{ or .Author "none"}}
Pipeline: {{ or .URL "none"}}
`,
			// Required:    true,
			Destination: &c.message,
		},
		&cli.BoolFlag{
			Name:        "colors",
			Usage:       "Want some color?",
			Destination: &c.colors,
		},
	}
}

func (c *Cli) Send(m Metadata) error {
	tmpl, err := template.New("message").Parse(c.message)
	if err != nil {
		return err
	}

	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	if c.colors {
		if m.ProjectTitle != "" {
			m.ProjectTitle = cyan(m.ProjectTitle)
		}

		if m.ProjectURL != "" {
			m.ProjectURL = yellow(m.ProjectURL)
		}

		if m.Branch != "" {
			m.Branch = yellow(m.Branch)
		}

		if m.CommitSHA != "" {
			m.CommitSHA = yellow(m.CommitSHA)
		}

		if m.Author != "" {
			m.Author = yellow(m.Author)
		}

		if m.URL != "" {
			m.URL = yellow(m.URL)
		}
	}

	err = tmpl.Execute(os.Stdout, m)
	if err != nil {
		return err
	}

	return nil
}
