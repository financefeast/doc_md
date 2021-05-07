---
title: Cashflow Examples
weight: 20
---

## Get Cashflow records
By sending a GET request to our [`/financial/cashflow]({{< relref "/api-documentation/api-v1/financial/financial-cashflow.md" >}}) endpoint you
will get cashflow financial records pulled from the companies annual report. Interim reports are not available. By default if no date or year
query parameters are passed to the endpoint the current years records will be returned.

{{< code-example exampleId="financial-cashflow" pathURL="/api-documentation/how-to/financial-cashflow/examples">}}

