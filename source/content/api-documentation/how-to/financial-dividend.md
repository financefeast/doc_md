---
title: Dividend Examples
weight: 40
---

## Get Balance Sheet records
By sending a GET request to our [`/financial/dividend]({{< relref "/api-documentation/api-v1/financial/financial-dividend.md" >}}) endpoint you
will get stock dividend payouts. By default, if no date or year query parameters are passed to the endpoint the current years records will be returned.

{{< code-example exampleId="financial-dividend" pathURL="/api-documentation/how-to/financial-dividend/examples">}}

