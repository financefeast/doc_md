from financefeast import Rest

client = Rest(token="your api token")
print(client.validate())

"""
This will return True if your token is valid, otherwise a 403 unauthorised will be returned.
"""