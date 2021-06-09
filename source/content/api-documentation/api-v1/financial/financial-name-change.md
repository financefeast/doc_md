---
title: Name Change
weight: 140
---

The Name Change endpoint returns company name change information. When a listed company changes its name its identifiers
such as symbol also change. Old and new information is listed in these records.

{{< rest-endpoint resource="financial-name-change" method="GET" path="/financial/name-change" >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}
{{< note >}} Only paid subscription plans can use this endpoint {{< /note >}}

## NameChange Entity

### Example
{{< rest-entity-example name="financialnamechange">}}

### Properties
{{< rest-entity-desc name="financialnamechange">}}

