---
title: Authentication
weight: 20
---

Before calling protected endpoints you will need to authenticate using your client id and client secret to obtain
an authorization token.

Find your client id and client secret [here](https://customer.financefeast.io). 

{{< rest-endpoint resource="authentication" method="GET" path="/oauth/login" >}}

## Authentication Entity

### Example
{{< rest-entity-example name="authentication">}}

### Properties
{{< rest-entity-desc name="authentication">}}

### Code Example

[Authentication]({{< relref "/api-documentation/how-to/authenticate.md" >}})

