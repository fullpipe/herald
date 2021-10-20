package target

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fullpipe/herald/place"
	"github.com/urfave/cli/v2"
)

type Grafana struct {
	host        string
	apiKey      string
	dashboardId int
	panelId     int
	message     string
	tags        cli.StringSlice
}

func (g *Grafana) Name() string {
	return "grafana"
}

func (g *Grafana) Usage() string {
	return "Add grafana anotations to your dashboards"
}

func (g *Grafana) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "host",
			Usage:       "Grafana api host, for example 'http://localhost:3000'",
			Destination: &g.host,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "api-key",
			Usage:       "Api key, editor permissions required",
			Destination: &g.apiKey,
			Required:    true,
		},
		&cli.StringSliceFlag{
			Name:        "tag",
			Aliases:     []string{"t"},
			Usage:       "Annotation tags",
			Destination: &g.tags,
		},
		&cli.StringFlag{
			Name:        "message",
			Aliases:     []string{"m"},
			Usage:       "Annotation description. Use metadata vars: {{.project}}, ...",
			Value:       `Project {{.project}} deployed, URL: {{ or .url "none"}}, Branch: {{ or .branch "none"}}, CommitSHA: {{ or .sha "none"}}, Author: {{ or .author "none"}}`,
			Destination: &g.message,
		},
		&cli.IntFlag{
			Name:        "dashboard",
			Usage:       "Dashboard ID",
			Destination: &g.dashboardId,
			DefaultText: "none",
		},
		&cli.IntFlag{
			Name:        "panel",
			Usage:       "Panel ID",
			Destination: &g.panelId,
			DefaultText: "none",
		},
	}
}

func (g *Grafana) Send(m place.Metadata) error {
	message, err := RenderMessage(g.message, m)
	if err != nil {
		return err
	}

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	postData := map[string]interface{}{
		"time": time.Now().UnixMilli(),
		"tags": g.tags.Value(),
		"text": message,
	}

	if g.dashboardId != 0 {
		postData["dashboardId"] = g.dashboardId
	}

	if g.panelId != 0 {
		postData["panelId"] = g.panelId
	}

	postBody, _ := json.Marshal(postData)

	req, err := http.NewRequest("POST", g.host+"/api/annotations", bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.apiKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		// TODO: unmarshal error from body {"message":"Unauthorized"}
		return errors.New(string(body))
	}

	return nil
}
