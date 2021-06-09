---
title: Bonus Issue
weight: 70
---

The Bonus endpoint returns stock bonus issue information. A bonus issue is an issue of shares of common stock to existing
shareholders instead of/or along with a dividend payment.

{{< rest-endpoint resource="financial-bonus" method="GET" path="/financial/bonus" >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}
{{< note >}} Only paid subscription plans can use this endpoint {{< /note >}}

## Bonus Entity

### Example
{{< rest-entity-example name="financialbonus">}}

### Properties
{{< rest-entity-desc name="financialbonus">}}

