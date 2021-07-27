from financefeast import Rest

client = Rest(client_id="your-client-id", client_secret="your-client-secret")
print(client.validate())

"""
This will return True if your token is valid, otherwise a 403 unauthorised will be returned.
"""