---
title: Stream API v1
weight: 10
---
# Stream API v1
Financefeast's stream API uses websockets to stream price data as soon as it is received from the exchange. For investors needing the absolute latest prices this is the best tool.

In order to start the stream with Financefeast Stream API, please sign up for an account
[here](https://identity.financefeast.io/account/signup).

Once you have signed up and have familiarized yourself with our API, please
check out our [python client](https://github.com/financefeast/python_client)
to begin integration with your own project.

## Authorization

### Stream API
Financefeast Stream API uses 20 bit authentication tokens that do not expire. The authentication token is then presented when you start the stream and then authorized. 
If the authentication code is valid, access is granted, otherwise a 403 Unauthorized error is returned.

The stream API uses [WebSocket Protocol (RFC 6455 & 7692)](https://datatracker.ietf.org/doc/html/rfc7692), so ensure your client is capable of that standard.

### Getting Your API Authentication Token

Make a note of your `api token`. These can be found [here](https://customer.financefeast.io/#creds). Paid plans allow for multiple token's to be issued which
allows you to stop reusing tokens between applications which decreases the risk if a token gets exposed.

## Domains
Financefeast's production stream API domain is `stream.financefeast.io`

## General Rules

### JSON
The stream API receives and sends JSON. Text strings are not supported.

### Time Format and Time Zone
All date time type inputs and outputs are serialized according to
[ISO8601](https://www.iso.org/iso-8601-date-and-time-format.html)
(more specifically [RFC3339](https://tools.ietf.org/html/rfc3339)).  The
communication does not assume a particular time zone, and this date time
serialization denominates the time offset of each value.

### Numbers
Decimal numbers are returned as floats or integers depending on the parameter. When sending integer or float values
to Financefeast API you can use strings or the primitive type.
