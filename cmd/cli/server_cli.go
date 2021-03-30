package main

import (
	"log"
	"os"

	"github.com/Liptonski/CLI_Web_Server/pkg/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Web Server CLI"
	app.Usage = "Runs html files on localhost"

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "runs specified html file on local server",

			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "specify file name as value",
				},
			},
			Action: func(c *cli.Context) error {
				server.Run_file()
				return nil
			},
		},
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
