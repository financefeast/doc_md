from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.eod(ticker='air.nz').data)