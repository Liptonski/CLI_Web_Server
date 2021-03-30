package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Web Server CLI"
	app.Usage = "Runs html files on localhost"

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "run --file <file_name>",

			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "file",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("Hello")
				return nil
			},
		},
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
