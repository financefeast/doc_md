from financefeast import Rest

client = Rest(client_id="your-client-id", client_secret="your-client-secret")
print(client.announcement(ticker="air.nz", data_from="2021-01-01").data)