---
title: Orderbook
weight: 50
---

The Orderbook endpoint returns the current level 2 market depth or orderbook for a stock. The orderbook is updated frequently
during trading hours on the exchange and should be considered as "real time".

The endpoint will, by default return a condensed orderbook showing collapsed price and quantity values. You can supply
a query parameter `condensed=false` to the endpoint to return the full orderbook.

{{< rest-endpoint resource="coredata-orderbook" method="GET" path="/data/orderbook" >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}
{{< note >}} Only subscription plans with Level 2 Market Data can use this endpoint {{< /note >}}

## Orderbook Entity

### Example Condensed (default)
{{< rest-entity-example name="coredataorderbookcondensed">}}

### Example Full
{{< rest-entity-example name="coredataorderbookfull">}}

### Properties
{{< rest-entity-desc name="coredataorderbook">}}

