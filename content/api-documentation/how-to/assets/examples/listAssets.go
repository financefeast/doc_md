package main

import (
	"github.com/financefeasthq/financefeast-trade-api-go/financefeast"
)

func init() {
	financefeast.SetBaseUrl("https://paper-api.financefeast.markets")
}

func main() {
	// Get a list of all active assets.
	status := "active"
	assets, err := financefeast.ListAssets(&status)
	if err != nil {
		panic(err)
	}

	// Filter the assets down to just those on NASDAQ.
	nasdaq_assets := []financefeast.Asset{}
	for _, asset := range assets {
		if asset.Exchange == "NASDAQ" {
			nasdaq_assets = append(nasdaq_assets, asset)
		}
	}
}