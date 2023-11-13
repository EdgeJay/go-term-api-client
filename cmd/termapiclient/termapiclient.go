package main

import (
	cli "github.com/EdgeJay/go-term-api-client/internal/cli"
	"github.com/EdgeJay/go-term-api-client/internal/termapiclient"
)

func main() {

	// load config and init cli
	cli := cli.NewRootApp(
		"termapiclient",
		"API client that runs in terminal, written in Go",
		`API client that makes API requests based on OpenAPI 3.x specifications. To get started, use

termapiclient run`,
	)

	cli.Configure()
	cli.AddCommand(
		"run",
		"Starts API client in terminal",
		`Starts API client in terminal.`,
		func(args []string) {
			// init terminal UI
			client := termapiclient.NewTermApiClient(nil)
			if err := client.Start(); err != nil {
				panic(err)
			}
		},
	)
	cli.Execute()
}
