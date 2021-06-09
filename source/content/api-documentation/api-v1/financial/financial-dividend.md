---
title: Dividends
weight: 50
---

The Dividend endpoint returns stock dividend information. From June 2021 the data is enhanced and shows various new fields
such as tax, currency code, exchange rate, period and comments. Prior to June 2021 only the 'dividend' field is present.

{{< rest-endpoint resource="financial-dividend" method="GET" path="/financial/dividend" >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}
{{< note >}} Only paid subscription plans can use this endpoint {{< /note >}}

## Dividend Entity

### Example
{{< rest-entity-example name="financialdividend">}}

### Properties
{{< rest-entity-desc name="financialdividend">}}

