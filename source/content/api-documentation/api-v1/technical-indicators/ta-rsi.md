---
title: Relative Strength Indicator
weight: 50
---

The Relative Strength Indicator endpoint returns price data along with a relative strength indicator
calculation. By default the lookback window the average is calculated on is 14 price points, but this can be set as a 
query parameter.

{{< rest-endpoint resource="ta-rsi" method="GET" path="/ta/rsi" >}}

## RSI Entity

### Example
{{< rest-entity-example name="tarsi">}}

### Properties
{{< rest-entity-desc name="tarsi">}}

