---
title: API v1
weight: 10
---
# API v1
Financefeast is a modern platform to obtain stock market data using a simple HTTP based API. Financefeast's
API allows easy retrieval of 20 minute delayed and real-time price and other financial information, such as company
financial records like cashflow and balance sheet data. Alternate data is also available for social media posts and
company announcements.

In order to start data retrieval with financefeast API, please sign up
[here](https://identity.financefeast.io/account/signup).

Once you have signed up and have familiarized yourself with our API, please
check out our [python client](https://github.com/financefeast/python_client)
to begin integration with your own project.

## Authorization

### API
Financefeast API uses 20 bit authentication tokens that do not expire. The authentication token is then presented to all protected endpoints and is then authorized. 
If the authentication code is valid, access is granted, otherwise a 403 Unauthorized error is returned.

### Integrated Applications
Financefeast API also allows Oauth2 authorization protocol that allows flexibility to developers writing applications that consume the Financefeast API. 
With oauth2 we can authorize customers in your behalf to access the API using their credentials which opens up many possibilities. 
Get in touch or read our guides on how you can develop your own Web application that uses oauth2 to authorize Finacefeast customers.

### Getting Your API Authentication Token

Make a note of your `api token`. These can be found [here](https://customer.financefeast.io/#creds). Paid plans allow for multiple token's to be issued which
allows you to stop reusing tokens between applications which decreases the risk if a token gets exposed.

### To Authorize to Protected Endpoints

To authorize to protected endpoints such as Core Data endpoints such as  End of Day or Intraday you will need to present an Authorization header with your api
authentication token. 
Calling a protected endpoint without providing a valid Authorization header will return a 403 Unauthorized error.

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
Object IDs in financefeast system uses UUID v4.  When making requests, the format
with dashes is accepted.

```
1d72e892-7336-4097-a762-7a9680111721
```

### Data Intervals and Data Retention
Financefeast API has 3 native time series intervals for its price data. These are:
| Interval | Parameter | Retention Policy
| -------- | ----------| ----------------
| 1 minute | 1m        | 1 year
| 1 hour   | 1h        | 3 years
| 1 day    | 1d        | 20 years

* When supplying an interval parameter to one of the API endpoints the data retention will be determined by the native intervals
in the table above. Ie using interval=1h will mean you will only be able to request 3 years of historical data.
* The retention policy is the minimum we will time series we will hold, but expect to retrain longer periods over time. 

Interval parameters can be supplied as:
| Interval | | | |
| -------  |-|-|-|
| weeks    | weeks | w | wks
| days     | days  | d | dys
| hours    | hours | h | hrs
| minutes  | minutes | m | mins
| seconds  | seconds | s | secs
