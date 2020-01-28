package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jschwinger23/aaec/ui/http"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:  "aaectl",
		Usage: "a simple command line client for aaecd",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "endpoint",
				Usage: `aaecd server netloc, such as "unix:///aaec.sock" or "http://localhost:2333"`,
			},
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:  "event",
				Usage: "events commands",
				Subcommands: []*cli.Command{
					&cli.Command{
						Name:  "create",
						Usage: "create new events",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "pkg",
								Usage: "package name",
							},
							&cli.StringFlag{
								Name:  "type",
								Usage: `event type, such as "bg" / "fg"`,
							},
						},
						Action: createEvents,
					},
				},
			},
			&cli.Command{
				Name:  "inst",
				Usage: "instruction commands",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "pkg",
						Usage: "package name",
					},
					&cli.StringSliceFlag{
						Name:  "extra",
						Usage: "instruction specific parameters, could set multiple times",
					},
				},
				Action: instruct,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}

	os.Exit(0)
}

func createEvents(c *cli.Context) (err error) {
	event := http.NewEvent(c.String("pkg"), c.String("type"))
	uri := NewURI(c.String("endpoint"))
	return uri.Compose("/events").Post(event)
}

func instruct(c *cli.Context) (err error) {
	extra := map[string]string{}
	for _, e := range c.StringSlice("extra") {
		parts := strings.Split(e, ",")
		extra[parts[0]] = parts[1]
	}

	instruction := http.NewInstruction(c.String("pkg"), extra)
	uri := NewURI(c.String("endpoint"))
	return uri.Compose("/inst").Post(instruction)
}
