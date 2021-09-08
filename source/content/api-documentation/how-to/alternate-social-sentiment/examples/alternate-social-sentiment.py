from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.social_sentiment(ticker='air.nz', data_from="2021-01-01").aggregation)