from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.ema(ticker='air.nz').data)