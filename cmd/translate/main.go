package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:      "gen-id",
				Usage:     "<file>",
				UsageText: "Generate IDs for each string in the given file",
				Action:    genIDCmd,
			},
			{
				Name:      "translate",
				Aliases:   []string{"tr"},
				Usage:     "<dir> <lang> <string>",
				UsageText: "Translate the provided string to the given language using the translations in the directory.",
				Action:    translateCmd,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
