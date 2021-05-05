---
title: Moving Average Convergence Divergence
weight: 40
---

The Moving Average Convergence Divergence endpoint returns price data along with a moving average convergence divergence
calculation. A signal line is calculated from a 9 exponential moving average. 26 price point intervals are required to 
produce the data points so for timeframes less than 26 data points the MACD will return 0.

{{< rest-endpoint resource="ta-macd" method="GET" path="/ta/macd" >}}

## MACD Entity

### Example
{{< rest-entity-example name="tamacd">}}

### Properties
{{< rest-entity-desc name="tamacd">}}

