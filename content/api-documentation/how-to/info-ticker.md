---
title: Ticker Examples
weight: 60
---

## Get a List of Tickers
By sending a GET request to our [`/info/ticker]({{< relref "/api-documentation/api-v1/info-ticker.md" >}}) endpoint you
will get a list of available tickers for the exchange passed as a query parameter. If no parameter is passed all tickers
on all exchanges will be returned.

{{< code-example exampleId="eod" pathURL="/api-documentation/how-to/info-ticker/examples">}}

