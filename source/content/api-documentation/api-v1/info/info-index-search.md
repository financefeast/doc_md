---
title: Index Search
weight: 100
---

The Index Search endpoint returns a list of available indices based on a search string.
The search string can be a partial or full string for an index_id, an index name or a uuid4.

{{< rest-endpoint resource="info-index-search" method="GET" path="/info/index/{string}" >}}

## Index Entity

### Example
{{< rest-entity-example name="infoindexsearch">}}

### Properties
{{< rest-entity-desc name="infoindexsearch">}}

