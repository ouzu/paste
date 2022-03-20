package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage: "share your files",
		Commands: []*cli.Command{
			{
				Name:      "upload",
				Aliases:   []string{"up", "u"},
				Usage:     "upload a file",
				ArgsUsage: "[file] (will read from stdin if no file specified)",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "filename after upload",
					},
				},
				Action: UploadHandler,
			},
			{
				Name:      "download",
				Aliases:   []string{"down", "d"},
				Usage:     "download a file",
				ArgsUsage: "[url]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "filename after download",
					},
				},
				Action: DownloadHandler,
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "print debug log",
				Value: false,
			},
			&cli.StringFlag{
				Name:  "server",
				Usage: "server address",
				Value: "paste.laze.today",
			},
		},
	}

	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
