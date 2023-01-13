package main

import (
	"hash/crc64"
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v2"
	"go.arsenm.dev/logger/log"
	"go.arsenm.dev/translate"
)

func genIDCmd(c *cli.Context) error {
	if c.NArg() < 1 {
		log.Fatal("This command requires one or more arguments").Send()
	}

	fl, err := os.Open(c.Args().First())
	if err != nil {
		log.Fatal("Error opening file").Err(err).Send()
	}

	var tr translate.Translations
	err = toml.NewDecoder(fl).Decode(&tr)
	if err != nil {
		return err
	}
	fl.Close()

	hash := crc64.New(crc64.MakeTable(crc64.ISO))

	for i, item := range tr.Items {
		hash.Reset()

		_, err := io.WriteString(hash, item.Value)
		if err != nil {
			return err
		}

		tr.Items[i].ID = hash.Sum64()
	}

	fl, err = os.Create(c.Args().First())
	if err != nil {
		log.Fatal("Error creating file").Err(err).Send()
	}

	err = toml.NewEncoder(fl).Encode(tr)
	if err != nil {
		log.Fatal("Error writing toml to file").Err(err).Send()
	}

	return nil
}
