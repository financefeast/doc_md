---
title: Exponential Moving Average
weight: 30
---

The Exponential Moving Average endpoint returns price data along with a exponential moving average calculation. By default
the lookback window the average is calculated on is 30 price points, but this can be set as a query parameter.

{{< rest-endpoint resource="ta-ema" method="GET" path="/ta/ep-ma" >}}

## EMA Entity

### Example
{{< rest-entity-example name="taema">}}

### Properties
{{< rest-entity-desc name="taema">}}

