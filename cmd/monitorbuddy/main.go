package main

import (
	"fmt"
	"log"

	"github.com/sstallion/go-hid"

	"leovanalphen/monitorbuddy/internal/app"
	"leovanalphen/monitorbuddy/internal/cli"
	"leovanalphen/monitorbuddy/internal/properties"
)

func main() {
	cli.Parse() // parse flags first for registry options etc.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Handle version early—no HID init needed
	if cli.FlagVersion() {
		// Keep the output simple and script-friendly
		fmt.Printf("mbuddy %s (commit %s, %s)\n", app.Version, app.Commit, app.Date)
		return
	}

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
