package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Liptonski/CLI_Web_Server/pkg/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Web Server CLI"
	app.Usage = "Runs html files on localhost"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Runs specified html file on local server",

			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "file",
					Usage:    "specify file name as value",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				file := os.Args[3]
				server.Run_file(file)
				return nil
			},
		},
		{
			Name:  "version",
			Usage: "Installed version of cli",
			Action: func(c *cli.Context) error {
				fmt.Println(app.Version)
				return nil
			},
		},
		// help is already implemented
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
