package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

type Slack struct{}

func (t *Slack) Send(m Metadata) error {
	webhook := ""

	// tmpl, err := template.New("").Parse("<!channel> {{.ProjectTitle}} deployed\nPipeline: {{.PipelineURL}}\n")
	// if err != nil {
	// 	panic(err)
	// }
	// err = tmpl.Execute(os.Stdout, sweaters)
	// if err != nil {
	// 	panic(err)
	// }

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

func (s *Slack) Flags() []cli.Flag {
	return []cli.Flag{
		// &cli.BoolFlag{
		// 	Name:        "colors2",
		// 	Usage:       "Want some color?",
		// 	Destination: &c.colors2,
		// },
		// &cli.BoolFlag{
		// 	Name:  "colors",
		// 	Usage: "Want some color?",
		// },
	}
}

func (c *Slack) Name() string {
	return "slack"
}
