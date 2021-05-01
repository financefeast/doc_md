---
title: API v1
weight: 10
---
# API v1
Financefeast is a modern platform to obtain stock market data using a simple HTTP based API. Financefeast's
API allows easy retrieval of 20 minute delayed and real-time price and fundamentals from multiple exchanges.

In order to start data retrieval with financefeast API, please sign up
[here](https://identity.financefeast.io/account/signup).

Once you have signed up and have familiarized yourself with our API, please
check out our [python client](https://github.com/financefeast/python_client)
to begin writing your own data visualization and processing application!

## Authorization and Authentication

Financefeast API uses the oauth2 authorization protocol and requires an initial authorization step to receive an authorization code. 
The authorization code is then presented to all protected endpoints, such as <code>Core Data</code> endpoints and is then authorized. 
If the authorization code is valid access is granted, otherwise a 401 Unauthorized error is returned.

Oauth2 authorization protocol allows flexibility to developers writing applications that consume the Financefeast API. 
With oauth2 we can authorize customers in your behalf to access the API using their credentials which opens up many possibilities. 
Get in touch or read our guides on how you can develop your own Web application that uses oauth2 to authorize Finacefeast customers.

### To Authenticate and Get an Authorization Code

First make a note of your client id and client secret. These can be found [here](https://customer.financefeast.io). 

[Authentication]({{< relref "./authentication.md" >}})

### To Authorize to Protected Endpoints

To authorize to protected endpoints such as Core Data: End of Day or Intraday you will need to present an Authorization header with your authorization code . 
Calling a protected endpoint without providing a valid Authorization header will return a 401 Unauthorized error.

[Authorization]({{< relref "./authorization.md" >}})

## Domains
Financefeast's production API domain is `api.financefeast.io`

## Rate Limit
There is a rate limit for the API requests.  When it is exceeded, the API
server returns error response with HTTP status code 429.  **The rate limit varies depending on the
subscription plan you are on.**

## General Rules

### Ticker symbols and uuid4
When using a ticker parameter on an endpoint you can use either:
* Symbol
* Uuid4

Either are accepted, with uuid4 being explict, where symbol could match from multiple exchanges if the [`exchange`] parameter is not passed
as well.

### Time Format and Time Zone
All date time type inputs and outputs are serialized according to
[ISO8601](https://www.iso.org/iso-8601-date-and-time-format.html)
(more specifically [RFC3339](https://tools.ietf.org/html/rfc3339)).  The
communication does not assume a particular time zone, and this date time
serialization denominates the time offset of each value.

### Numbers
Decimal numbers are returned as floats or integers depending on the parameter. When sending integer or float values
to Financefeast API you can use strings or the primitive type.

### IDs
Object ID in financefeast system uses UUID v4.  When making requests, the format
with dashes is accepted.

```
1d72e892-7336-4097-a762-7a9680111721
```
