---
title: Cashflow
weight: 20
---

The Cashflow endpoint returns cashflow financial information from the companies annual report. Only annual report data
is available, not interim and is updated daily.

{{< rest-endpoint resource="financial-cashflow" method="GET" path="/financial/cashflow" >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}
{{< note >}} Only paid subscription plans can use this endpoint {{< /note >}}

## Cashflow Entity

### Example
{{< rest-entity-example name="financialcashflow">}}

### Properties
{{< rest-entity-desc name="financialcashflow">}}

