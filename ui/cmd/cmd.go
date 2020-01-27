package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jschwinger23/aaec/ui"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
)

type Event = struct {
	ID       string
	CreateAt int64
	Type     string
	Content  []byte
}

func main() {
	if err := func() (err error) {
		if err != nil {
			return
		}
		app := cli.App{
			Name:  "aaectl",
			Usage: "a simple command line client for aaecd",
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
		return app.Run(os.Args)
	}(); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}

	os.Exit(0)
}

func createEvents(c *cli.Context) (err error) {
	application, err := ui.GetApplication()
	if err != nil {
		return
	}

	types := strings.Split(c.String("types"), ",")
	contents := c.StringSlice("contents")
	if len(types) != len(contents) {
		return errors.New("types not coincide with contents")
	}

	events := []Event{}
	for i := 0; i < len(types); i++ {
		events = append(events, Event{
			ID:       uuid.New().String(),
			CreateAt: time.Now().Unix(),
			Type:     types[i],
			Content:  []byte(contents[i]),
		})
		fmt.Printf("ui.cmd generated event: %v", events[i])
	}

	status, err := application.CreateEvents(events)
	if status != "ok" {
		err = errors.Wrapf(err, "status exception: %s", status)
	}
	return
}
