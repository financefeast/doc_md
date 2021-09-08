from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.exchange_status().data)