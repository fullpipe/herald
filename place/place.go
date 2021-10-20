package place

import "github.com/urfave/cli/v2"

type Place interface {
	Name() string
	Usage() string
	Metadata() (Metadata, error)
	Flags() []cli.Flag
}

type Metadata struct {
	Project string
	URL     string
	Branch  string
	SHA     string
	Author  string
}
