package main

import (
	"fmt"
	"github.com/financefeasthq/financefeast-trade-api-go/financefeast"
)

func init() {
	financefeast.SetBaseUrl("https://paper-api.financefeast.markets")
}

func main() {
	// Check if AAPL is tradable on the financefeast platform.
	asset, err := financefeast.GetAsset("AAPL")
	if err != nil {
		fmt.Println("Asset not found for AAPL.")
	} else if asset.Tradable {
		fmt.Println("We can trade AAPL.")
	}
}