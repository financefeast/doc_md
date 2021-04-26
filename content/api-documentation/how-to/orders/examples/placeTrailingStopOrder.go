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

	// Submit a trailing stop order to sell 1 share of Apple at a
    // trailing stop of
	symbol = "AAPL"
	financefeast.PlaceOrder(financefeast.PlaceOrderRequest{
		AssetKey: &symbol,
		Qty: decimal.NewFromFloat(1),
		Side: financefeast.Sell,
		Type: financefeast.TrailingStop,
		StopPrice: decimal.NewFromFloat(1.00),  // stop price will be hwm - 1.00$
		TimeInForce: financefeast.Day,
	})

	// Alternatively, you could use trail_percent:
	symbol = "AAPL"
	financefeast.PlaceOrder(financefeast.PlaceOrderRequest{
		AssetKey: &symbol,
		Qty: decimal.NewFromFloat(1),
		Side: financefeast.Sell,
		Type: financefeast.TrailingStop,
		TrailPercent: decimal.NewFromFloat(1.0),  // stop price will be hwm*0.99
		TimeInForce: financefeast.Day,
	})
}