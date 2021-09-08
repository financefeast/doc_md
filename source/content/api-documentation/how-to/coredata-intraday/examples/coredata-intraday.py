from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.intraday(ticker='air.nz').data)