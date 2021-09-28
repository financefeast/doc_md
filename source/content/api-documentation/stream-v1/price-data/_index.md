---
title: Price Data
weight: 20
---
# Price Data

## Authentication

Upon establishing a socket connection to the stream API you need to authenticate by sending a JSON payload to the API

```json
{"type": "authenticate",
    "data": {
        "token": "your API token"
    }}
```
If your authentication is successful you will receive a response
```json
{'type': 'authenticated', 
  'data': {'message': 'You have been authenticated to the Stream api. Data stream now starts'}
}
```
Failure to authenticate can be either; You have no access to the Stream API or you are using all of your available streams.

From this point on price data will stream to your client in JSON format. You should design your client to expect disconnection at any time and
auto reconnect.

## Price Data

All symbols will be sent to your client when received from the exchange. There is currently no option to pre-filter symbols. They will be in JSON format
```json
{'type': 'price', 'data': {'date': '2021-09-29T09:33:30', 'price': 2.61, 'volume': 5000, 'symbol': 'SUM020.NZ'}}
{'type': 'price', 'data': {'date': '2021-09-29T09:34:45', 'price': 3.0, 'volume': 55000, 'symbol': 'WIA080.NZ'}}
{'type': 'price', 'data': {'date': '2021-09-29T09:34:45', 'price': 3.0, 'volume': 20000, 'symbol': 'WIA080.NZ'}}
{'type': 'price', 'data': {'date': '2021-09-29T09:34:50', 'price': 2.65, 'volume': 25000, 'symbol': 'SUM030.NZ'}}
{'type': 'price', 'data': {'date': '2021-09-29T09:35:53', 'price': 3.0424, 'volume': 181123, 'symbol': 'NZB.NZ'}}
{'type': 'price', 'data': {'date': '2021-09-29T09:36:17', 'price': 99.9, 'volume': 5000, 'symbol': 'KCFHA.NZ'}}
{'type': 'price', 'data': {'date': '2021-09-29T09:36:28', 'price': 2.9491, 'volume': 408798, 'symbol': 'NZC.NZ'}}
{'type': 'price', 'data': {'date': '2021-09-29T09:41:27', 'price': 99.8, 'volume': 50000, 'symbol': 'ANBHB.NZ'}}
```

### Data Structure

| Parameter | Type | Description
| --------- | ---- | -----------
| type | string | The type of data received, can be either 'price', 'authenticated' or 'error'
| data | dictionary | The data dictionary
| date | string | The date of the price update
| price | float | The price update
| volume | integer | The volume of units in the price update
| symbol | string | The symbol for the price update

All data sent from the stream API is in JSON format.

## Logoff
To logoff from the stream API simply drop the connection. 

## Errors

### No Stream Access
```json
{'type': 'error', 'data': {'message': 'Your subscription plan has no access to the Stream API. Please upgrade your plan'}}
```
Upgrade your plan to get access to the stream API

### Unauthenticated
```json
{'type': 'error', 'data': {'message': 'Token not authorised. Please check your token and try again'}}
```
Check your API token

### Too many streams
```json
{'type': 'error', 'data': {'message': 'You are at your stream limit of X and cannot establish further streams until you remove one'}}
```

## Examples Using Financefeast Client Library
The [Financefeast](https://pypi.org/project/financefeast/) client library makes it easy to consume the stream API and is recommended.

### Installation
Use pip to install the library
```shell
pip install financefeast
```

### Usage
Create an instance of the Stream class, supplying your token as a parameter and optionally a callback method to handle the data
```python
from financefeast.stream import Stream

def on_data(stream, data):
    print(data)

client = Stream(token='tk_oey965IWImNlIXQMcpQ2', on_data=on_data)
client.connect()
```
