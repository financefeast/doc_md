---
title: Try Out and Test the Financefeast API
weight: 10
description: How to test the Financefeast API without coding.
---

## How to get started and test the Financefeast API
Want to get stock data with Financefeast, but you aren’t fully sure how the Financefeast API works yet? Here’s a step by step tutorial on testing the Financefeast API by manually using the Postman interface.

## This tutorial explains how to:
* Install Postman
* Get our Financefeast postman collection
* Find your API client credentials
* Configure the Postman collection
* Authenticate to the Financefeast API  
* Get stock data through API

### Install Postman
If you don't already have Postman installed, [Get it here](https://app.getpostman.com/app/download/win64).
We recommend the stand alone version of Postman, rather than the Chrome browser plugin due to its age.

### Get our Postman Collection
[Available on Github](https://github.com/financefeast/postman_collections)

```commandline
git clone https://github.com/financefeast/postman_collections.git
```

### Get Your API Client Credentials
[Credentials are here](https://customer.financefeast.io)

Copy your client id and client secret. You will need to enter these into environment variables in the postman collection

### Configure Postman
Under "Environments", select "Prod" and you should see a list of variables. 
Enter your client id into CLIENT_ID and your client secret into CLIENT_SECRET.

{{< safe-html >}}
<center><img src="./environment.png" width="80%"></center>
{{< /safe-html >}}

### Authenticate
Under "Collections", run "Login" which will attempt to authenticate your client credentials. 
If successful you will see a Bearer token. This will automatically be saved as an environment 
variable "TOKEN" and used to authenticate all other requests to end points.

### Get Stock Data
Once you are authenticated, under "Collections", run "EOD". You should see data for the pre-defined ticker. 
You can change the ticker, and the date by changing the query params.