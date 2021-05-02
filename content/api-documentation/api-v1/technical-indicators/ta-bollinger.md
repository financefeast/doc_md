---
title: Bollinger Band
weight: 70
---

The Bollinger Band endpoint returns price data along with a bollinger band calculation. 
By default, the Bollinger Band lookback window is 20 price points, but you can specify any lookback window via a
query parameter.

{{< rest-endpoint resource="ta-bollinger" method="GET" path="/ta/bollinger" >}}

## Bollinger Entity

### Example
{{< rest-entity-example name="tabollinger">}}

### Properties
{{< rest-entity-desc name="tabollinger">}}

