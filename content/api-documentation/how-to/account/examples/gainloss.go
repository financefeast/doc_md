package main

import (
	"fmt"
	"log"

	"github.com/financefeasthq/financefeast-trade-api-go/financefeast"
)

func main() {
	financefeast.SetBaseUrl("https://paper-api.financefeast.markets")

	// Get account information.
	account, err := financefeast.GetAccount()
	if err != nil {
		log.Fatalln(err)
	}

	// Calculate the difference between current balance and balance at the last market close.
	balanceChange := account.Equity.Sub(account.LastEquity)

	fmt.Println("Today's portfolio balance change:", balanceChange)
}
