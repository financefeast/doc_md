from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.sma(ticker='air.nz').data)