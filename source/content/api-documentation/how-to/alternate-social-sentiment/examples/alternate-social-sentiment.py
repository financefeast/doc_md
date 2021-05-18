from financefeast import FinanceFeast, Environments

client = FinanceFeast(client_id="your-client-id", client_secret="your-client-secret")
print(client.social_media(ticker='air.nz', data_from="2021-01-01"))