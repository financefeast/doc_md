---
title: Income
weight: 30
---

The Income endpoint returns income financial information from the companies annual report. Only annual report data
is available, not interim and is updated daily.

{{< rest-endpoint resource="financial-income" method="GET" path="/financial/income >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}
{{< note >}} Only paid subscription plans can use this endpoint {{< /note >}}

## Income Entity

### Example
{{< rest-entity-example name="financialincome">}}

### Properties
{{< rest-entity-desc name="financialincome">}}

