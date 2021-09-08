from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.bollinger(ticker='air.nz').data)