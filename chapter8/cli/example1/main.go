package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Create new app
	app := cli.NewApp()

	// Add flags with three arguments
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "stranger",
			Usage: "Your wonderful name",
		},
		cli.IntFlag{
			Name:  "age",
			Value: 0,
			Usage: "Your graceful age",
		},
	}

	// This function parses and brings data in cli.Context struct
	app.Action = func(c *cli.Context) error {
		// c.String, c.Int looks for value of given flag
		log.Printf("Hello %s (%d Years), Welcome to the command line world", c.String("name"), c.Int("age"))
		return nil
	}

	// Pass os.Args to cli app to parse content
	app.Run(os.Args)
}
