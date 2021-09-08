---
title: Using the Financefeast Python Client Library
weight: 20
description: Get familiar with using the python client library to authenticate, pull data and graph
---

## Using the Financefeast Python Client Library
Get familiar with using the python client library to authenticate, pull data and graph. In this guide we'll show
you how easy it is to pull data and visualize it.

## This tutorial explains how to:
* Install the Financefeast python client library
* Authenticate to the API
* Pull historical end of day data
* Graph


### Get Your API Token
[Credentials are here](https://customer.financefeast.io/#creds)

Copy API authentication token. You will need to enter this into a parameter when creating an instance of Financefeast in the script

### Install the Financefeast python client library
* Ensure you have [python 3.8](https://www.python.org/downloads/release/python-380) installed.
* From a shell type:
```commandline
pip install financefeast
```

### Authenticate to the API
* Open your favorite IDE or text editor and create a new project or file
* Import the following:
```python
from financefeast import Rest, Environments
import pandas as pd
import matplotlib.pyplot as plt
```
This imports the Financefeast client library, pandas for data manipulation and matplotlib for graphing.

{{< safe-html >}}
<br>
{{< /safe-html >}}

Now authenticate to the API
```python
client = Rest(token="YOUR API TOKEN")
```
This should print a message like this:
```python
INFO:ff_client:API environment set as prod
INFO:ff_client:Client successfully authorized to API
```

### Pull End of Day Data
We will get EOD data for 'air.nz' for the year to date:
```python
result = client.eod(ticker='air.nz', date_from="2021-01-01")
```
If you print the 'data' variable you should see something like this (I've abbreviated it for display):
```python
[{'timestamp': '2021-01-07T00:00:00', 'close': 1.79, 'open': 1.8, 'high': 1.8, 'low': 1.785, 'volume': 291549.0},..., ....]
```
Now we'll put the data into a Pandas DataFrame so that we can manipulate the data and plot it:
```python
df = pd.DataFrame(result.data)
```
Note that we call the the `data` property of the `result` object to access the price data.

### Graph the Data
Now that we have the price data in a Pandas DataFrame we will visualize it. First we need to set the index on the DataFrame
as type datetime:
```python
df['date'] = pd.to_datetime(df['timestamp'])
df.set_index('date', inplace=True)
```
Now we can get Pandas to plot the graph using the 'close' column as the data source
```python
gph = df['close'].plot()
plt.show()
```
You should now be looking at a graph like this
{{< safe-html >}}
<center><img src="./graph.png" width="80%"></center>
{{< /safe-html >}}

### Final Code
Your final code should look like this:
```python
from financefeast import Rest, Environments
import pandas as pd
import matplotlib.pyplot as plt

client = Rest(token="YOUR API TOKEN")
result = client.eod(ticker='air.nz', date_from="2021-01-01")
df = pd.DataFrame(result.data)
df['date'] = pd.to_datetime(df['timestamp'])
df.set_index('date', inplace=True)
gph = df['close'].plot()
plt.show()
```
