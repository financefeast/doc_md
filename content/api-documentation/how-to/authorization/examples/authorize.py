from financefeast import FinanceFeast, Environments

client = FinanceFeast(client_id="your-client-id", client_secret="your-client-secret", environment=Environments.prod.value)
print(client.validate())

"""
This will return True if your token is valid, otherwise a 403 unauthorised will be returned.
"""