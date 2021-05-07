---
title: Balance Sheet
weight: 40
---

The Balance endpoint returns balance sheet financial information from the companies annual report. Only annual report data
is available, not interim and is updated daily.

{{< rest-endpoint resource="financial-balance" method="GET" path="/financial/balance >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}
{{< note >}} Only paid subscription plans can use this endpoint {{< /note >}}

## Balance Entity

### Example
{{< rest-entity-example name="financialbalance">}}

### Properties
{{< rest-entity-desc name="financialbalance">}}

