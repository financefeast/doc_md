---
title: Last Examples
weight: 180
---

## Get Last Prices
By sending a GET request to our [`/data/last`]({{< relref "/api-documentation/api-v1/price-data/coredata-last.md" >}}) endpoint you
will get the last available price data for a ticker. This can be called intraday and on non trading days. This is guaranteed to return
a single price data record, unlike the [`/data/eod`] endpoint which, if called without date parameters on a non trading day will not return data. 

{{< code-example exampleId="coredata-last" pathURL="/api-documentation/how-to/coredata-last/examples">}}

