from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.ticker(exchange='nzx').data)