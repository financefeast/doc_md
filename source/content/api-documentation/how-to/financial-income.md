---
title: Income Examples
weight: 30
---

## Get Cashflow records
By sending a GET request to our [`/financial/income`]({{< relref "/api-documentation/api-v1/financial/financial-income.md" >}}) endpoint you
will get income financial records pulled from the companies annual report. Interim reports are not available. By default, if no date or year
query parameters are passed to the endpoint the current years records will be returned.

{{< code-example exampleId="financial-income" pathURL="/api-documentation/how-to/financial-income/examples">}}

