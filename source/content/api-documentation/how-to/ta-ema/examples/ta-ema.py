from financefeast import Rest

client = Rest(client_id="your-client-id", client_secret="your-client-secret")
print(client.ema(ticker='air.nz').data)