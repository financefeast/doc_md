---
title: Using the Financefeast Python Client Library to Graph Sentiment
weight: 30
description: Graph stock sentiment against stock price
---

## Visualizing Stock Sentiment vs Stock Price
Does sentiment around a particular stock affect the stock price? Its an interesting question and something we can
graph up and test. In this tutorial we will pull stock sentiment and stock price and graph.

## This tutorial explains how to:
* Install the Financefeast python client library
* Authenticate to the API
* Pull historical end of day data
* Pull sentiment data  
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
from financefeast import Rest
import matplotlib.pyplot as plt
from datetime import datetime
```
This imports the Financefeast client library and matplotlib for graphing.

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

### Pull End of Day and Sentiment Data
We will get EOD and sentiment data for 'air.nz' for the last few months:
```python
sentiments = client.social_sentiment(ticker="air.nz", date_from="2021-06-01")
price = client.eod(ticker='air.nz', date_from="2021-06-01")
```
If you print the 'data' variable you should see something like this (I've abbreviated it for display):
```python
[{'timestamp': '2021-01-07T00:00:00', 'close': 1.79, 'open': 1.8, 'high': 1.8, 'low': 1.785, 'volume': 291549.0},..., ....]
```
We need to extract the dictionaries holding the postive and negative sentiment:
```python
positive_sentiment = sentiments.aggregation[0]['days']
negative_sentiment = sentiments.aggregation[1]['days']
```
And we need to turn the dictionaries for sentiment into seperate lists for plotting the x and y axes for both sentiment and price
```python
p_x = list(map(lambda x:datetime.strptime(x['created'], "%Y-%m-%dT%H:%M:%S.%fZ"),positive_sentiment)) # note this is to the microsecond
p_y = list(map(lambda x:x['count'],positive_sentiment))

n_x = list(map(lambda x:datetime.strptime(x['created'], "%Y-%m-%dT%H:%M:%S.%fZ"),negative_sentiment)) # note this is to the microsecond
n_y = list(map(lambda x:x['count'],negative_sentiment))

c_x = list(map(lambda x:datetime.strptime(x['timestamp'], "%Y-%m-%dT%H:%M:%S"),price.data))   # note this is to the second
c_y = list(map(lambda x:x['close'],price.data))
```

### Graph the Data
Now that we have x and y axes data for sentiment and price we can plot these.
```python
fig, ax= plt.subplots(figsize=(20, 10))
fig.tight_layout()
ax2=ax.twinx()
ax3=ax.twinx()
ax.plot(p_x, p_y, color="green", label="pos")
ax2.plot(n_x, n_y, color="red", label="neg")
ax3.plot(c_x, c_y, color="blue", label="close")
ax.set_title('AIR.NZ Close Price vs Social Sentiment')
plt.show()
```

You should now be looking at a graph like this
{{< safe-html >}}
<center><img src="./sentiment_graph.png" width="80%"></center>
{{< /safe-html >}}

### Final Code
Your final code should look like this:
```python
from financefeast import Rest
import matplotlib.pyplot as plt
from datetime import datetime

"""
Pull the data from Financefeast
NOTE: Remember to put in your client_id and client_secret
"""
client = Rest(token="YOUR API TOKEN")
sentiments = client.social_sentiment(ticker="air.nz", date_from="2021-06-01")
price = client.eod(ticker='air.nz', date_from="2021-06-01")

"""
We need to extract the postive and negative sentiment counts
"""
positive_sentiment = sentiments.aggregation[0]['days']
negative_sentiment = sentiments.aggregation[1]['days']

"""
Then using map pull the timestamp and count values into a list ready for plotting
"""
p_x = list(map(lambda x:datetime.strptime(x['created'], "%Y-%m-%dT%H:%M:%S.%fZ"),positive_sentiment)) # note this is to the microsecond
p_y = list(map(lambda x:x['count'],positive_sentiment))

n_x = list(map(lambda x:datetime.strptime(x['created'], "%Y-%m-%dT%H:%M:%S.%fZ"),negative_sentiment)) # note this is to the microsecond
n_y = list(map(lambda x:x['count'],negative_sentiment))

c_x = list(map(lambda x:datetime.strptime(x['timestamp'], "%Y-%m-%dT%H:%M:%S"),price.data))   # note this is to the second
c_y = list(map(lambda x:x['close'],price.data))

"""
Create the plots
"""

fig, ax= plt.subplots(figsize=(20, 10))
fig.tight_layout()
ax2=ax.twinx()
ax3=ax.twinx()
ax.plot(p_x, p_y, color="green", label="pos")
ax2.plot(n_x, n_y, color="red", label="neg")
ax3.plot(c_x, c_y, color="blue", label="close")
ax.set_title('AIR.NZ Close Price vs Social Sentiment')
plt.show()
```

## Closing Thoughts
Does sentiment affect price? No, for this data range it does not appear so. This is a simple example to show how easy it is
to pull data out of the Financefeast API. Many factors affect stock price and looking at a single variable in isolation will be
unlikely to show results.

If you would like to see the complete Google Colab notebook for this please [visit here](https://colab.research.google.com/drive/1BhNqhpNrE4-mtNELtOdc6RuY3kuAUo_S?usp=sharing)
