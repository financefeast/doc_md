---
title: Intraday Examples
weight: 40
---

## Get Intraday Prices
By sending a GET request to our [`/data/intraday`]({{< relref "/api-documentation/api-v1/coredata-intraday.md" >}}) endpoint you
will get intraday price data for the ticker requested. Prices are updated at least every minute, but depending on your
subscription plan they maybe delayed by 20 minutes. [Check your subscription plan](https://customer.financefeast.io/#sub) for
details.

{{< code-example exampleId="eod" pathURL="/api-documentation/how-to/coredata-intraday/examples">}}

