package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var name string
	app := cli.NewApp()
	app.Name = "echo-cli"
	app.Version = "0.0.1"
	app.Usage = "This echo whatever you provide to the app as args"
	app.ArgsUsage = "echo app [arg 1] [args2] ..."
	app.Authors = []cli.Author{
		{
			Name:  "Richard",
			Email: "ricdon41@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "Input name",
			Hidden:      false,
			Value:       "Richard",
			Destination: &name,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q", name)
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "details",
			Usage: "echo the full details of the user",
			Action: func(c *cli.Context) error {
				fmt.Printf("Echoing the details")
				return nil
			},
		},
	}

	app.Run(os.Args)
}
