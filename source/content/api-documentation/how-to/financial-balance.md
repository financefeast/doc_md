---
title: Balance Examples
weight: 30
---

## Get Balance Sheet records
By sending a GET request to our [`/financial/balance`]({{< relref "/api-documentation/api-v1/financial/financial-balance.md" >}}) endpoint you
will get balance sheet financial records pulled from the companies annual report. Interim reports are not available. By default, if no date or year
query parameters are passed to the endpoint the current years records will be returned.

{{< code-example exampleId="financial-balance" pathURL="/api-documentation/how-to/financial-balance/examples">}}

