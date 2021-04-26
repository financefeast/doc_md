package main

import (
	"github.com/financefeasthq/financefeast-trade-api-go/financefeast"
	"github.com/shopspring/decimal"
)

func init() {
	financefeast.SetBaseUrl("https://paper-api.financefeast.markets")
}

func main() {
	// Submit a market order to buy 1 share of Apple at market price
	symbol := "AAPL"
	financefeast.PlaceOrder(financefeast.PlaceOrderRequest{
		AssetKey: &symbol,
		Qty: decimal.NewFromFloat(1),
		Side: financefeast.Buy,
		Type: financefeast.Market,
		TimeInForce: financefeast.Day,
	})

	// Submit a limit order to attempt to sell 1 share of AMD at a
	// particular price ($20.50) when the market opens
	symbol = "AMD"
	financefeast.PlaceOrder(financefeast.PlaceOrderRequest{
		AssetKey: &symbol,
		Qty: decimal.NewFromFloat(1),
		Side: financefeast.Sell,
		Type: financefeast.Limit,
		TimeInForce: financefeast.OPG,
		LimitPrice: decimal.NewFromFloat(20.50),
	})
}