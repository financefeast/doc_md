---
title: Conversion
weight: 100
---

The Conversion endpoint returns convertible stock conversion information. Convertible stock is exchanged for a set
number of common stock usually by a set date, or some other trigger.

{{< rest-endpoint resource="financial-conversion" method="GET" path="/financial/conversion" >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}
{{< note >}} Only paid subscription plans can use this endpoint {{< /note >}}

## Conversion Entity

### Example
{{< rest-entity-example name="financialconversion">}}

### Properties
{{< rest-entity-desc name="financialconversion">}}

