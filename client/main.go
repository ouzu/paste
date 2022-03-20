package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage: "share your files",
		Commands: []*cli.Command{
			{
				Name:    "upload",
				Aliases: []string{"up", "u"},
				Usage:   "upload a file",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}},
					&cli.StringFlag{Name: "stdin", Aliases: []string{"s"}},
				},
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args())
					fmt.Println("added task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "download",
				Aliases: []string{"down", "d"},
				Usage:   "download a file",
				Action: func(c *cli.Context) error {
					fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
		},
	}

	app.EnableBashCompletion = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
