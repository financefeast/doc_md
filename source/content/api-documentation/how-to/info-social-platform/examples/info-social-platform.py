from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.social_platform().data)