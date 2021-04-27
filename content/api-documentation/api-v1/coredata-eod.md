---
title: End of Day
weight: 30
---

The EOD endpoint returns end-of-day stock price data for high, low, open, close, volume. This can be called at any time, even intra-day and 
will return the last price received to date.

{{< rest-endpoint resource="coredata-eod" method="GET" path="/data/eod" >}}

{{</* info title="Authorization" */>}}
Authorization required for this endpoint
{{</* /note */>}}

## EOD Entity

### Example
{{< rest-entity-example name="coredataeod">}}

### Properties
{{< rest-entity-desc name="coredataeod">}}

