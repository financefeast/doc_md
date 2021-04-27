---
title: Intraday
weight: 30
---

The Intraday endpoint returns end-of-day stock price data for high, low, open, close and volume. Prices are updated at least
every minute and when called this endpoint will return the latest price available, plus historical depending on the datetime range
supplied.

{{< rest-endpoint resource="core-data-intraday" method="GET" path="/data/intraday" >}}

<aside class="info">
Authorization required for this endpoint
</aside>

## EOD Entity

### Example
{{< rest-entity-example name="coredataintraday>}}

### Properties
{{< rest-entity-desc name="coredataintraday">}}

