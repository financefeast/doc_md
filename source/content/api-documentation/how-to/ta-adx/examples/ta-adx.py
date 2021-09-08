from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.adx(ticker='air.nz').data)