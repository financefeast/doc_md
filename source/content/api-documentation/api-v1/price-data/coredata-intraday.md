---
title: Intraday
weight: 30
---

The Intraday endpoint returns intraday stock price data for high, low, open, close and volume. Prices are updated at least
every minute and when called this endpoint will return the latest price available, plus historical price data depending on the datetime range
supplied.

{{< rest-endpoint resource="coredata-intraday" method="GET" path="/data/intraday" >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}

## EOD Entity

### Example
{{< rest-entity-example name="coredataintraday">}}

### Properties
{{< rest-entity-desc name="coredataintraday">}}

