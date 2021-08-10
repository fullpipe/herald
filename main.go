package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "herald",
		Usage: "notify someone from somewhere",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "message",
				Aliases: []string{"m"},
				Usage:   "Message to send `MESSAGE`",
			},
		},
		Action: func(c *cli.Context) error {
			place, err := GetPlace(c.Args().First())
			if err != nil {
				return err
			}

			target, err := GetTarget(c.Args().Get(1))
			if err != nil {
				return err
			}

			meta, err := place.GetMetadata()
			if err != nil {
				return err
			}

			message, err := GetMessage(meta, c.String("message"))
			if err != nil {
				return err
			}

			return target.Send(message)
		},
	}

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

type Place interface {
	GetMetadata() (Metadata, error)
}

type Target interface {
	Send(m Message) error
}

type Gitlab struct{}

func (g *Gitlab) GetMetadata() (Metadata, error) {
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

type Slack struct{}

func (t *Slack) Send(m Message) error {
	webhook := "https://hooks.slack.com/services/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"

	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	attachment := slack.Attachment{
		Color:      "good",
		AuthorName: "Herald",
		//AuthorSubname: "github.com",
		//AuthorLink:    "https://github.com/slack-go/slack",
		//AuthorIcon:    "https://avatars2.githubusercontent.com/u/652790",
		Text: "<!channel> All text in Slack uses the same system of escaping: chat messages, direct messages, file comments, etc. :smile:\nSee <https://api.slack.com/docs/message-formatting#linking_to_channels_and_users>",
		Ts:   json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	err := slack.PostWebhook(webhook, &msg)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("sending to slack:", m)
	return nil
}
