package target

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/fullpipe/herald/place"
	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

type Slack struct {
	webhook string
	message string
	color   string
}

func (s *Slack) Name() string {
	return "slack"
}

func (s *Slack) Usage() string {
	return "Send notifications to slack"
}

func (s *Slack) Send(m place.Metadata) error {
	message, err := RenderMessage(s.message, m)
	if err != nil {
		return err
	}

	fields := []slack.AttachmentField{}
	if m.Project != "" {
		fields = append(fields, slack.AttachmentField{Title: "Project", Value: m.Project, Short: true})
	}
	if m.URL != "" {
		fields = append(fields, slack.AttachmentField{Title: "URL", Value: m.URL, Short: true})
	}
	if m.Branch != "" {
		fields = append(fields, slack.AttachmentField{Title: "Branch", Value: m.Branch, Short: true})
	}
	if m.SHA != "" {
		fields = append(fields, slack.AttachmentField{Title: "SHA", Value: m.SHA, Short: true})
	}
	if m.Author != "" {
		fields = append(fields, slack.AttachmentField{Title: "Author", Value: m.Author, Short: true})
	}

	attachment := slack.Attachment{
		Color:      "good",
		AuthorName: "Herald",
		Ts:         json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
		Fields:     fields,
	}

	msg := slack.WebhookMessage{
		Text:        message,
		Attachments: []slack.Attachment{attachment},
	}

	err = slack.PostWebhook(s.webhook, &msg)
	if err != nil {
		return err
	}

	return nil
}

func (s *Slack) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "webhook",
			Usage:       "Incoming Webhooks `URL`",
			Destination: &s.webhook,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "color",
			Usage:       "Message color, good|warning|danger|#439FE0",
			Destination: &s.color,
			Value:       "good",
		},
		&cli.StringFlag{
			Name:        "message",
			Aliases:     []string{"m"},
			Usage:       "Message temlate. Use metadata vars: {{.project}}, ...",
			Value:       `<!channel> Project {{.project}} deployed`,
			Destination: &s.message,
		},
	}
}
