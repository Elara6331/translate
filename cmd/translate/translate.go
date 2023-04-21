package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"go.elara.ws/logger/log"
	"go.elara.ws/translate"
	"golang.org/x/text/language"
)

func translateCmd(c *cli.Context) error {
	if c.NArg() < 3 {
		log.Fatal("This command requires three or more arguments").Send()
	}

	t, err := translate.NewFromDir(c.Args().First())
	if err != nil {
		log.Fatal("Error creating new translator").Err(err).Send()
	}

	tag, err := language.Parse(c.Args().Get(1))
	if err != nil {
		log.Fatal("Error parsing language").Err(err).Send()
	}

	fmt.Println(t.TranslateTo(c.Args().Get(2), tag))
	return nil
}
