---
title: Authorization
weight: 25
---

To authorize to protected endpoints you must pass an 'Bearer Authorization' http header containing a token that was recently 
returned from a successful authenticate call. Tokens expire, therefore detecting "403 unauthorised" errors when using
long running tasks is "best practice" to ensure expected application behaviour.

{{< safe-html >}}
<br>
{{< /safe-html >}}

To take the hassle out of authentication and authorization you should use our client libraries to take care of these
tasks for you, so you can concentrate on developing your features. [SDK for your language]({{< relref "/api-documentation/client-sdk/_index.md" >}})

## Authorization Entity

You can find more information on how to form the header [here]({{< relref "/api-documentation/how-to/authorization.md" >}})

### Example
{{< rest-entity-example name="authorization">}}

