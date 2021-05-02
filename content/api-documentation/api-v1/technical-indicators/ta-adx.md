---
title: Average Directional Index
weight: 60
---

The Average Directional Index endpoint returns price data along with an average directional index
calculation. The ADX calculation requires two lookback windows. The first defaults to 5 price points and
the second to 15. These are the default recommended, but can be changed using query parameters.

{{< rest-endpoint resource="ta-adx" method="GET" path="/ta/adx" >}}

## ADX Entity

### Example
{{< rest-entity-example name="taadx">}}

### Properties
{{< rest-entity-desc name="taadx">}}

