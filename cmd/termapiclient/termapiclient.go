package main

import "github.com/EdgeJay/go-term-api-client/internal/termapiclient"

func main() {
	client := termapiclient.NewTermApiClient(nil)
	if err := client.Start(); err != nil {
		panic(err)
	}
}
