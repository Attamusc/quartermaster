package main

import (
    "github.com/attamusc/quartermaster/quartermaster"
    "github.com/codegangsta/cli"
    "os"
)

func main() {
    app := cli.NewApp()
    app.Name = "quartermaster"
    app.Usage = "Manage and install your client side dependencies with class"
    app.Version = "0.0.1"

    app.Commands = []cli.Command{
        {
            Name:  "install",
            Usage: "Explicitly state a dependency to install",
            Action: func(c *cli.Context) {
                quartermaster.Install()
            },
        },
        {
            Name:  "bundle",
            Usage: "Install dependencies defined in quartermaster.json file",
            Action: func(c *cli.Context) {
                quartermaster.Bundle()
            },
        },
    }

    app.Run(os.Args)
}
