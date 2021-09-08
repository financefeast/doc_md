from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.stochastic(ticker='air.nz').data)