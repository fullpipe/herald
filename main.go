package main

import (
	"log"
	"os"

	"github.com/fullpipe/herald/place"
	"github.com/fullpipe/herald/target"
	"github.com/urfave/cli/v2"
)

func main() {
	places := []place.Place{&place.Nowhere{}, &place.Git{}, &place.Gitlab{}}
	targets := []target.Target{&target.Cli{}, &target.Grafana{}, &target.Slack{}}

	app := cli.NewApp()

	app.Usage = "Notify someone from some place about your success"
	app.UsageText = "herald PLACE [place options] TARGET [target options]"

	for _, pl := range places {
		subCommands := []*cli.Command{}

		for _, t := range targets {
			subCommands = append(subCommands, &cli.Command{
				Name:  t.Name(),
				Usage: t.Usage(),
				Flags: t.Flags(),
				Action: func(pl place.Place, t target.Target) func(c *cli.Context) error {
					return func(c *cli.Context) error {
						meta, err := pl.Metadata()
						if err != nil {
							return err
						}

						return t.Send(meta)
					}
				}(pl, t),
			})
		}

		app.Commands = append(app.Commands, &cli.Command{
			Name:        pl.Name(),
			Usage:       pl.Usage(),
			Flags:       pl.Flags(),
			Subcommands: subCommands,
		})
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
