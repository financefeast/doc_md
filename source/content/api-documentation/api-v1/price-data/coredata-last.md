---
title: Last
weight: 60
---

The Last endpoint returns the last recorded price data for high, low, open, close, volume. This can be called at any time, even intra-day and 
will return the last price received to date. This is updated hourly and if called on non trading days it will then return the last trading days
end of day price data.

{{< rest-endpoint resource="coredata-last" method="GET" path="/data/last" >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}

## Last Entity

### Example
{{< rest-entity-example name="coredatalast">}}

### Properties
{{< rest-entity-desc name="coredatalast">}}

