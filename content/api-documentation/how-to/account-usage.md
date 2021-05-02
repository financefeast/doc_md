---
title: Usage Examples
weight: 90
---

## List and Count Endpoints Called
By sending a GET request to our [`/account/usage']({{< relref "/api-documentation/api-v1/account-usage.md" >}}) endpoint you
will get a list of all the authenticated endpoints you have called for the date range passed. By default, a month will be returned.

{{< safe-html >}}
<br>
{{< /safe-html >}}

The list will contain each endpoint called along with a data field containing each day and a count for that day.

{{< code-example exampleId="account-usage" pathURL="/api-documentation/how-to/account-usage/examples">}}

