package target

import (
	"bytes"
	"text/template"

	"github.com/fullpipe/herald/place"
	"github.com/urfave/cli/v2"
)

type Target interface {
	Name() string
	Usage() string
	Send(meta place.Metadata) error
	Flags() []cli.Flag
}

// Project: {{ or .project "none"}}
// URL: {{ or .url "none"}}
// Branch: {{ or .branch "none"}}
// CommitSHA: {{ or .sha "none"}}
// Author: {{ or .author "none"}}
func RenderMessage(message string, meta place.Metadata) (string, error) {
	tmpl, err := template.New("message").Parse(message)
	if err != nil {
		return "", err
	}

	out := bytes.NewBufferString("")

	err = tmpl.Execute(out, map[string]string{
		"project": meta.Project,
		"url":     meta.URL,
		"branch":  meta.Branch,
		"sha":     meta.SHA,
		"author":  meta.Author,
	})
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
