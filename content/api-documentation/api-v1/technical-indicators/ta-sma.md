---
title: Simple Moving Average
weight: 20
---

The Simple Moving Average endpoint returns price data along with a simple moving average calculation. By default
the lookback window the average is calculated on is 30 price points, but this can be set as a query parameter.

{{< rest-endpoint resource="ta-sma" method="GET" path="/ta/sm-ma" >}}

## SMA Entity

### Example
{{< rest-entity-example name="tasma">}}

### Properties
{{< rest-entity-desc name="tasma">}}

