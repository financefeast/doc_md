---
title: Ticker Search Examples
weight: 70
---

## Get a List of Tickers Based on a Search String
By sending a GET request to our [`/info/ticker/{string}`]({{< relref "/api-documentation/api-v1/info-ticker-search.md" >}}) endpoint you
will get a list of available tickers based on a search sring for the exchange passed as a query parameter. If no parameter is passed all tickers
on all exchanges will be returned.

{{< code-example exampleId="info-ticker-search" pathURL="/api-documentation/how-to/info-ticker-search/examples">}}

