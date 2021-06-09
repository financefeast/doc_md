---
title: Change of Class
weight: 90
---

The Change of Class endpoint returns information when a class change for stock occurs. This occurs when changes between
either common stock or preferred stock change classes, eg units of common stock converted to preferred stock and vice
versa.

{{< rest-endpoint resource="financial-change-of-class" method="GET" path="/financial/change-of-class" >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}
{{< note >}} Only paid subscription plans can use this endpoint {{< /note >}}

## Class Entity

### Example
{{< rest-entity-example name="financialchangeofclass">}}

### Properties
{{< rest-entity-desc name="financialchangeofclass">}}

