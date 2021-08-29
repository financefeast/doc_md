---
title: Company Announcements
weight: 30
---

The Company Announcements endpoint returns company announcements as advised through NZX exchange as part of the rules of listing on the exchange.
The announcements are unstructured text in raw format as returned from the exchange. The text will contain `\r` return character 
codes indicating an end of line. 

{{< rest-endpoint resource="alternate-company-announcements" method="GET" path="/alternate/announcement" >}}

{{< note >}} Authorization is required for this endpoint {{< /note >}}
{{< note >}} Only paid subscription plans can use this endpoint {{< /note >}}

## Announcement Entity

### Example
{{< rest-entity-example name="alternateannouncement">}}

### Properties
{{< rest-entity-desc name="alternateannouncement">}}

