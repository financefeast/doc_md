from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.cpi(data_from="1990-01-01").data)