from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.last(ticker='air.nz').data)