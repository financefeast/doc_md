---
title: Authentication
weight: 20
---

Before calling protected endpoints you will need to authenticate using your client id and client secret to obtain
an authorization token.

{{< rest-endpoint resource="authentication" method="GET" path="/oauth/login" >}}

## Authentication Entity

### Example
{{< rest-entity-example name="authentication">}}

### Properties
{{< rest-entity-desc name="authentication">}}

### Code Example

[Authentication]({{< relref "/how-to/authentication.md" >}})