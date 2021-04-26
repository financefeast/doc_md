package main

import (
	"github.com/financefeasthq/financefeast-trade-api-go/financefeast"
)

func init() {
	financefeast.SetBaseUrl("https://paper-api.financefeast.markets")
}

func main() {
	// Get the last 100 of our closed orders
	status := "closed"
	limit := 100
	nested := true // show nested multi-leg orders
	closed_orders, err := financefeast.ListOrders(&status, nil, &limit, &nested)
	if err != nil {
		panic(err)
	}

	// Get only the closed orders for a particular stock
	var aaplOrders []financefeast.Order
	for _, order := range closed_orders {
		if order.Symbol == "AAPL" {
			aaplOrders = append(aaplOrders, order)
		}
	}
}