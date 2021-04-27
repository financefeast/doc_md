---
title: EOD Examples
weight: 30
---

## Get EOD Prices
By sending a GET request to our [`/data/eod`]({{< relref "/api-documentation/api-v1/coredata-eod.md" >}}) endpoint you
will get end of day price data for the ticker requested. The latest price is updated intraday so if you specify 'todays'
date you will get the latest price available. Historical dates will return the closing price.

{{< code-example exampleId="eod" pathURL="/api-documentation/how-to/coredata-eod/examples">}}

