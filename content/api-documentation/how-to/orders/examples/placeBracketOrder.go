package main

import (
    "github.com/financefeasthq/financefeast-trade-api-go/financefeast"
    "github.com/financefeasthq/financefeast-trade-api-go/common"
    "github.com/shopspring/decimal"
)

func init() {
    API_KEY := "YOUR_API_KEY_HERE"
    API_SECRET := "YOUR_API_SECRET_HERE"
    BASE_URL := "https://paper-api.financefeast.markets"
    // Check for environment variables
    if common.Credentials().ID == "" {
        os.Setenv(common.EnvApiKeyID, API_KEY)
    }
    if common.Credentials().Secret == "" {
        os.Setenv(common.EnvApiSecretKey, API_SECRET)
    }
    financefeast.SetBaseUrl(BASE_URL)
}

func main() {
    // Submit a limit order to buy 1 share of Apple and add
    // StopLoss and TakeProfit orders.
    client := financefeast.NewClient(common.Credentials())
    symbol := "AAPL"
    tpp := decimal.NewFromFloat(320.)
    spp := decimal.NewFromFloat(317.)
    limit := decimal.NewFromFloat(318.)
    tp := &financefeast.TakeProfit{LimitPrice: &tpp}
    sl := &financefeast.StopLoss{
        LimitPrice: nil,
        StopPrice:  &spp,
    }
    req := financefeast.PlaceOrderRequest{
        AccountID:   common.Credentials().ID,
        AssetKey:    &symbol,
        Qty:         decimal.New(1, 0),
        Side:        financefeast.Buy,
        LimitPrice:  &limit,
        TimeInForce: financefeast.GTC,
        Type:        financefeast.Limit,
        OrderClass:  financefeast.Bracket,
        TakeProfit:  tp,
        StopLoss:    sl,
    }
    order, err := client.PlaceOrder(req)
    fmt.Println(order)
    fmt.Println(err)

    // We could buy a position and just add a stop loss (OTO Orders)
    spp := decimal.NewFromFloat(317.)
    limit := decimal.NewFromFloat(318.)
    sl := &financefeast.StopLoss{
        StopPrice:  &spp,
    }
    req := financefeast.PlaceOrderRequest{
        AccountID:   common.Credentials().ID,
        AssetKey:    &symbol,
        Qty:         decimal.New(1, 0),
        Side:        financefeast.Buy,
        LimitPrice:  &limit,
        TimeInForce: financefeast.GTC,
        Type:        financefeast.Limit,
        OrderClass:  financefeast.Oto,
        StopLoss:    sl,
    }
    order, err := client.PlaceOrder(req)
    fmt.Println(order)
    fmt.Println(err)

    // We could split it to 2 orders. first buy a stock,
    // and then add the stop/profit prices (OCO Orders)
    limit := decimal.NewFromFloat(318.)
    req := financefeast.PlaceOrderRequest{
        AccountID:   common.Credentials().ID,
        AssetKey:    &symbol,
        Qty:         decimal.New(1, 0),
        Side:        financefeast.Buy,
        LimitPrice:  &limit,
        TimeInForce: financefeast.GTC,
        Type:        financefeast.Limit,
        OrderClass:  financefeast.Simple,
    }
    order, err := client.PlaceOrder(req)
    fmt.Println(order)
    fmt.Println(err)

    // wait for it to complete and then
    tpp := decimal.NewFromFloat(320.)
    spp := decimal.NewFromFloat(317.)
    tp := &financefeast.TakeProfit{LimitPrice: &tpp}
    sl := &financefeast.StopLoss{
        LimitPrice: nil,
        StopPrice:  &spp,
    }
    req := financefeast.PlaceOrderRequest{
        AccountID:   common.Credentials().ID,
        AssetKey:    &symbol,
        Qty:         decimal.New(1, 0),
        Side:        financefeast.Sell,
        TimeInForce: financefeast.GTC,
        Type:        financefeast.Limit,
        OrderClass:  financefeast.Oco,
        TakeProfit:  tp,
        StopLoss:    sl,
    }
    order, err := client.PlaceOrder(req)
    fmt.Println(order)
    fmt.Println(err)
}
