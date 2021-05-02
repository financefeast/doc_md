---
title: Stochastic Oscillator
weight: 80
---

The Stochastic Oscillator endpoint returns price data along with a stochastic oscillator calculation. 
By default, the Stochastic Oscillator lookback window is 14 price points, but you can specify any lookback window via a
query parameter. The calculation also requires a simple moving average lookback window and by default that is 3 price points.

{{< rest-endpoint resource="ta-stochastic" method="GET" path="/ta/stochastic" >}}

## Stochastic Entity

### Example
{{< rest-entity-example name="tastochastic">}}

### Properties
{{< rest-entity-desc name="tastochastic">}}

