from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.balance(ticker='air.nz', year=2020).data)