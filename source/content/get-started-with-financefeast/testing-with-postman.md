---
title: Test Financefeast API with Postman
weight: 10
description: How to test the Financefeast API without coding using Postman.
---

## Test Financefeast API with Postman
Want to get stock data with Financefeast, but you aren’t fully sure how the Financefeast API works yet? Here’s a step by step tutorial on testing the Financefeast API by manually using the Postman interface.

## This tutorial explains how to:
* Install Postman
* Get our Financefeast postman collection
* Find your API token
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

### Get Your API Token
[Credentials are here](https://customer.financefeast.io/#creds)

Copy your API token. You will need to enter these into environment variable in the postman collection

### Configure Postman
Under "Environments", select "Prod" and you should see a list of variables. 
Enter your API token into the TOKEN value field.

{{< safe-html >}}
<center><img src="./environment.png" width="80%"></center>
{{< /safe-html >}}

{{< note >}} Make sure the 'Environment' is set to 'Prod' {{< /note >}}

### Get Stock Data
Once you have entered your TOKEN into the TOKEN environment variable, under "Collections", run "EOD". You should see data for the pre-defined ticker. 
You can change the ticker, and the date by changing the query params.

{{< safe-html >}}
<center><img src="./data.png" width="80%"></center>
{{< /safe-html >}}