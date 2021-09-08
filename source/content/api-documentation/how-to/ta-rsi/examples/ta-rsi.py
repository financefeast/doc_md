from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.rsi(ticker='air.nz').data)