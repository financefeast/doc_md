package main

import (
	"fmt"

	"github.com/financefeasthq/financefeast-trade-api-go/financefeast"
    "github.com/financefeasthq/financefeast-trade-api-go/stream"
)

func init() {
	financefeast.SetBaseUrl("https://paper-api.financefeast.markets")
}

func main() {
	if err := stream.Register(financefeast.TradeUpdates, tradeHandler); err != nil {
        panic(err)
    }
}

func tradeHandler(msg interface{}) {
	tradeupdate := msg.(financefeast.TradeUpdate)
	if tradeupdate.Order.ClientOrderID == "my_client_order_id" {
		fmt.Printf("Update for {}. Event: {}.\n", tradeupdate.Order.ClientOrderID, tradeupdate.Event)
	}
}