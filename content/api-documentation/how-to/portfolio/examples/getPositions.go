package main

import (
	"fmt"
	"github.com/financefeasthq/financefeast-trade-api-go/financefeast"
)

func init() {
	financefeast.SetBaseUrl("https://paper-api.financefeast.markets")
}

func main() {
	// Get our position in AAPL.
	aapl_position, err := financefeast.GetPosition("AAPL")
	if err != nil {
		fmt.Println("No AAPL position.")
	} else {
		fmt.Printf("AAPL position: %v shares.\n", aapl_position.Qty)
	}

	// Get a list of all of our positions.
	positions, err := financefeast.ListPositions()
	if err != nil {
		fmt.Println("No positions found.")
	} else {
		// Print the quantity of shares for each position.
		for _, position := range positions {
			fmt.Printf("%v shares in %s", position.Qty, position.Symbol)
		}
	}
}