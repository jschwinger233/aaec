package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/imroc/req"
	"github.com/jschwinger23/aaec/ui/http"
	"github.com/pkg/errors"
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
								Name:  "types",
								Usage: `event types, separated by comma, such as "adb,timer,adb"`,
							},
							&cli.StringSliceFlag{
								Name:  "contents",
								Usage: `event contents, could set multiple times`,
							},
						},
						Action: createEvents,
					},
				},
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
	types := strings.Split(c.String("types"), ",")
	contents := c.StringSlice("contents")
	if len(types) != len(contents) {
		return errors.New("types not coincide with contents")
	}

	events := http.Events{}
	for i := 0; i < len(types); i++ {
		event := http.NewEvent(types[i], []byte(contents[i]))
		events.AddEvent(event)
		fmt.Printf("ui.cmd generated event: %v", event)
	}

	uri := c.String("endpoint") + "/events"
	r, err := req.Post(uri, req.BodyJSON(&events))
	if err != nil {
		return errors.Wrap(err, "ui.http.client failed to post events")
	}
	resp := r.Response()
	if resp.StatusCode > 300 {
		return errors.Errorf("ui.http.client http status exception: %v", resp.StatusCode)
	}

	return nil
}
