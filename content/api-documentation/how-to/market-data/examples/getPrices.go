package main

import (
	"fmt"
	"github.com/financefeasthq/financefeast-trade-api-go/financefeast"
)

func init() {
	financefeast.SetBaseUrl("https://paper-api.financefeast.markets")
}

func main() {
	// Get daily price data for AAPL over the last 5 trading days.
	barCount := 5
	bars, err := financefeast.GetSymbolBars("AAPL", financefeast.ListBarParams{
		Timeframe: "day",
		Limit: &barCount,
	})
	if err != nil {
		panic(err)
	}

	// See how much AAPL moved in that timeframe.
	startPrice := bars[0].Open
	endPrice := bars[len(bars)-1].Close
	percentChange := (endPrice - startPrice) / startPrice * 100
	fmt.Printf("AAPL moved %v%% over the last 5 days.", percentChange)
}