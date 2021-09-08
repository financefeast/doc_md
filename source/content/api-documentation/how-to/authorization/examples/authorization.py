from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.validate())

"""
This will return True if your token is valid, otherwise a 403 unauthorised will be returned.
"""