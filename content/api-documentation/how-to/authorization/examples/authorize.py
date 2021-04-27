from financefeast import FinanceFeast, Environments

client = FinanceFeast(client_id="your-client-id", client_secret="your-client-secret", environment=Environments.prod.value)
print(client.eod(ticker="air.nz"))