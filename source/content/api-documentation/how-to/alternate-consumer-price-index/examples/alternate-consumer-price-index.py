from financefeast import Rest

client = Rest(client_id="your-client-id", client_secret="your-client-secret")
print(client.cpi(data_from="1990-01-01").data)