---
title: Authorization
weight: 14
---

To authorize to protected endpoints you must pass an 'Bearer Authorization' http header containing your API authentication token which
can be obtained from your [API Dashboard]("https://customer.financefeast.io/#creds"). Tokens do not expire so care should be taken to
ensure they are not leaked. You can regenerate or delete your tokens from your API Dashboard at any time. Paid subscription plans allow
for multiple tokens.

Tokens are 20 bit length keys prefixed with either `tk_` for test regions or `pk_` for production. You should ensure you pass the correct
token to the correct region. They are not interchangeable and passing an incorrect token prefix will return a 403 Unauthorised.

{{< safe-html >}}
<br>
{{< /safe-html >}}

To take the hassle out of authentication and authorization you should use our client libraries to take care of these
tasks for you, so you can concentrate on developing your features. [SDK for your language]({{< relref "/api-documentation/client-sdk/_index.md" >}})

## Authorization Entity

You can find more information on how to form the header [here]({{< relref "/api-documentation/how-to/authorization.md" >}})

### Example
{{< rest-entity-example name="authorization">}}

### Code Example

[Authorization]({{< relref "/api-documentation/how-to/authorization.md" >}})