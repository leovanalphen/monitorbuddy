package main

import (
	"log"

	"github.com/sstallion/go-hid"

	"leovanalphen/monitorbuddy/internal/app"
	"leovanalphen/monitorbuddy/internal/cli"
	"leovanalphen/monitorbuddy/internal/properties"
)

func main() {
	cli.Parse() // parse flags first for registry options etc.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}
	defer hid.Exit()

	properties.BuildRegistry(cli.FlagIncludeGB())

	if cli.FlagHelpProps() {
		properties.PrintPropsTable()
		return
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
