from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.macd(ticker='air.nz').data)