package place

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

type Git struct {
	path string
}

func (g *Git) Name() string {
	return "git"
}

func (g *Git) Usage() string {
	return ""
}

func (g *Git) Metadata() (Metadata, error) {
	m := Metadata{}

	title, err := g.cmd("git", "rev-parse", "--show-toplevel")
	if err != nil {
		return m, err
	}
	m.Project = filepath.Base(title)

	branch, err := g.cmd("git", "branch", "--show-current")
	if err != nil {
		return m, err
	}
	m.Branch = branch

	sha, err := g.cmd("git", "rev-parse", "HEAD")
	if err != nil {
		return m, err
	}
	m.SHA = sha

	author, err := g.cmd("git", "show", "-s", "--format=%an <%ae>", "HEAD")
	if err != nil {
		return m, err
	}
	m.Author = author

	url, err := g.cmd("git", "remote", "get-url", "origin")
	if err != nil {
		return m, err
	}
	m.URL = url

	return m, nil
}

func (g *Git) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "path",
			Usage:       "path to your git repository",
			Value:       ".",
			Destination: &g.path,
		},
	}
}

func (g *Git) cmd(args ...string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = g.path
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}
