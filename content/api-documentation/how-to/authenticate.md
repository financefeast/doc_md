---
title: Authenticate Examples
weight: 10
---

## Authenticate
By sending a GET request to our [`/oauth/login`]({{< relref "/api-documentation/api-v1/authentication.md" >}}) endpoint
you receive an authorization token which is required to be passed to protected endpoints as an 'Authorization' http header.

You can get your client credentials [here]('https://customer.financefeast.io')

{{< code-example exampleId="authentication" pathURL="/api-documentation/how-to/authentication/examples">}}

